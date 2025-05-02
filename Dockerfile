# Use Go base image
FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .
RUN go build -o notes-api

EXPOSE 8080
CMD ["./notes-api"]
