FROM golang:1.21.11-alpine

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY *.go .
COPY . .

RUN go build -o service .

CMD ["./service"]