FROM golang:1.22-alpine as builder
WORKDIR /app
COPY . .
RUN go mod tidy

RUN go build -o ./arctic-wolf-risk-manager cmd/main.go

FROM alpine:latest
WORKDIR /code
COPY --from=builder /app/arctic-wolf-risk-manager ./arctic-wolf-risk-manager
CMD ["./arctic-wolf-risk-manager"]