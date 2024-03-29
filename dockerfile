FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download

# EXPOSE $PORT
EXPOSE 8080

CMD ["go", "run", "main.go", "server"]
