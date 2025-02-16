FROM golang:1.24.0-alpine

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY *.go .
COPY . .

RUN go build -o service .

CMD ["./service"]