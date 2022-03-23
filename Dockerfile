FROM golang:1.16-alpine

WORKDIR /app-market

COPY . .

RUN go build -o api-market

EXPOSE 8080

CMD ./api-market