# DEV
FROM golang:alpine3.20 AS build

WORKDIR /app
COPY . /app
RUN go mod download
RUN go build -o main main.go

# PROD
FROM alpine:3.20

WORKDIR /app
COPY --from=build /app/main /app/main
ENV PORT 3000
EXPOSE 3000
CMD ["./main"]