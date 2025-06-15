FROM golang:1.25rc1-alpine

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY *.go .
COPY . .

RUN go build -o service .

CMD ["./service"]