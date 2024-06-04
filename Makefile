.PHONY: postgres inmemory

postgres:
	docker-compose --profile postgresql up -d --build

inmemory:
	docker-compose --profile inmemory up -d --build

clean:
	docker-compose --profile postgresql --profile inmemory down --volumes

test:
	go test -race ./... -tags unit -coverprofile .testCoverage.out
	go tool cover -html=.testCoverage.out
	rm .testCoverage.out