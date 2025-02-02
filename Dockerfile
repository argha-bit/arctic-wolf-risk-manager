FROM golang:1.22.3 as builder
WORKDIR /app
COPY go.mod ./
RUN go mod tidy
COPY . .
RUN go build -o arctic-wolf-risk-manager cmd/main.go
EXPOSE 8080
CMD ["./arctic-wolf-risk-manager"]