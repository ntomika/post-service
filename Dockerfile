FROM golang:1.22-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main cmd/main.go

RUN apk add --no-cache --update git

EXPOSE 8080

CMD ["./main"]