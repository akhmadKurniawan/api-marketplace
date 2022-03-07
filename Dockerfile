FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go build -o api-marketing

EXPOSE 8080

CMD ./api-marketing