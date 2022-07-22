FROM golang:1.16-alpine
WORKDIR /app
COPY main.go ./
EXPOSE 8080
CMD go run main.go