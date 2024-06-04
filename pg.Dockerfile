FROM golang:1.22-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main cmd/main.go

RUN apk --no-cache add postgresql-client

COPY migrations/init.sql /docker-entrypoint-initdb.d/

COPY entrypoint.sh /app/entrypoint.sh

COPY app.env /app/app.env

RUN apk add --no-cache --update git

EXPOSE 8080

ENTRYPOINT ["sh", "/app/entrypoint.sh"]