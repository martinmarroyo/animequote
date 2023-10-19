FROM golang:1.20.10-alpine3.17

WORKDIR /app

COPY /app .

RUN go build .

ENTRYPOINT ["./animequote"]