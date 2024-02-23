FROM golang:1.21 AS builder
WORKDIR /go-calculator
COPY . .
RUN go build -o  /go-calculator/bin/app ./cmd/app/main.go

FROM alpine
COPY --from=builder /go-calculator/bin /go-calculator
CMD ["go-calculator/app"]