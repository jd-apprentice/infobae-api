# DEV
FROM golang:alpine3.20 AS build

WORKDIR /app
COPY . /app
RUN go mod download
RUN GOARCH=amd64 go build -o ./bin/InfobaeAPI ./src/main.go

# PROD
FROM alpine:3.20

WORKDIR /app
COPY --from=build /app/bin/InfobaeAPI /app/bin/InfobaeAPI
ENV GIN_MODE=release
ENV PORT=3000
EXPOSE 3000
ENTRYPOINT ["/app/bin/InfobaeAPI"]
