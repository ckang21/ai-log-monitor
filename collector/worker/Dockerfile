# worker/Dockerfile
FROM golang:1.24.1

WORKDIR /app

COPY . .

RUN go build -o worker main.go

CMD ["./worker"]
