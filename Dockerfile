FROM golang:1.19-alpine

WORKDIR /app

COPY . .

RUN go get -u ./...

RUN cd src && go build -o /server

EXPOSE 4000

CMD ["/server"]
