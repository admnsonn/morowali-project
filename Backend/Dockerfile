# Build stage
FROM golang:latest as build

WORKDIR /app
COPY . .

RUN go build -o main .

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy only the necessary files
COPY --from=build /app/main .

# Debugging: Print the contents of the /app directory
RUN ls -la

EXPOSE 7071

CMD ["./main"]
