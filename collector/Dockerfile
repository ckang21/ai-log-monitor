# collector/Dockerfile
FROM golang:1.24.1

WORKDIR /app

COPY . .

RUN go build -o collector main.go queue.go

CMD ["./collector"]
