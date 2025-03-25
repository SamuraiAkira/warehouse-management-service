FROM golang:1.24-alpine as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /warehouse-service ./cmd/app

FROM alpine:latest
COPY --from=builder /warehouse-service /warehouse-service
EXPOSE 8080
CMD ["/warehouse-service"]