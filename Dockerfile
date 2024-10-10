FROM golang:1.23.2-alpine3.20
COPY src/* /src
WORKDIR /src

RUN go build

ENTRYPOINT ["go", "run ./cmd/"]