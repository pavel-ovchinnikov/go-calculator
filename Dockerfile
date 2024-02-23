FROM golang:1.21 AS builder
WORKDIR /go-calculator
COPY . .
RUN go build -o ./bin/app ./cmd/app/main.go
CMD ["/go-calculator/bin/app"]