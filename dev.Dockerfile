FROM golang:1.23-alpine

WORKDIR /flowgo

RUN go install github.com/air-verse/air@latest

COPY ./flowgo/go.* .

RUN go mod download

COPY ./flowgo .

EXPOSE 8000

CMD [ "air", "-c", ".air.toml"]