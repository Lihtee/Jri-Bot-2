FROM golang:1.23.2-alpine3.20
COPY src/* /src
WORKDIR /bot

RUN go build -C ../src -o ../bot/bot

ENTRYPOINT ["/bot/bot"]