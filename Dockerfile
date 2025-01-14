FROM golang:1.23-alpine AS builder

WORKDIR /flowgo

COPY ./flowgo/go.* .

RUN go mod download

COPY ./flowgo .

RUN go build -o flowgo

FROM alpine:3.20.3 AS compiled

WORKDIR /root

COPY --from=builder /flowgo/flowgo .

EXPOSE 8000

ENTRYPOINT ["./flowgo"]