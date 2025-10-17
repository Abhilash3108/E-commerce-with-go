FROM golang:1.25.1-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ecommerce cmd/main.go
CMD ["./ecommerce"]
EXPOSE 8000
