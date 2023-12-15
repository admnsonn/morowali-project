# Use the official Golang base image
FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o main .
EXPOSE 7071

CMD ["./main"]
