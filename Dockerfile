FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -mod vendor -o main .

CMD ["./main"]
