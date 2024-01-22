# Build stage
FROM golang:latest as build

WORKDIR /app
COPY . .

RUN go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/main .

RUN ls -la

EXPOSE 8080

CMD ["./main"]
