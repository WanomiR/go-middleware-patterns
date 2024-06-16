FROM golang:1.19-alpine

WORKDIR /app

COPY . .

RUN go mod init github.com/wanomir/go-middleware-patterns

RUN go get -u github.com/joho/godotenv

RUN go build -o server

CMD ["server"]
