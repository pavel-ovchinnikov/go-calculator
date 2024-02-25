FROM golang:1.21 AS builder
WORKDIR /go-calculator
COPY . .
RUN go get ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/app ./cmd/app

FROM alpine
COPY --from=builder /go-calculator/bin/app /go-calculator/app
COPY --from=builder /go-calculator/config/config.yaml /go-calculator/config.yaml
CMD ["/go-calculator/app"]