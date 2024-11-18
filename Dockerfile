FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY ./app/go.* .

RUN go mod download

COPY ./app .

RUN go build -o flowgo

FROM alpine:3.20.3 AS compiled

WORKDIR /root

COPY --from=builder /app/flowgo .

EXPOSE 8000

ENTRYPOINT ["./flowgo"]