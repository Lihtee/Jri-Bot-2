FROM golang:1.23.0-alpine3.20

WORKDIR /src
COPY src .

RUN go build -o bot

ENTRYPOINT ["./bot"]