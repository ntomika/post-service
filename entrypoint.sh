#!/bin/sh

# Функция для проверки доступности базы данных
check_postgres() {
    until PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -c '\q' > /dev/null 2>&1; do
        echo "Postgres is unavailable - sleeping"
        sleep 10
    done

    echo "Postgres is up"
}

# Ожидание готовности базы данных
check_postgres

# Выполнение миграций
PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -f /docker-entrypoint-initdb.d/init.sql

./main