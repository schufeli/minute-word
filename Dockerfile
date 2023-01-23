FROM golang:1.19.5-alpine

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY *.go .

RUN go build -o main .

CMD ["/main"]