FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go build -o cake-store-api

EXPOSE 8001

CMD ./cake-store-api