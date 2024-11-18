FROM golang:1.23-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY ./app/go.* .

RUN go mod download

COPY ./app .

EXPOSE 8000

CMD [ "air", "-c", ".air.toml"]