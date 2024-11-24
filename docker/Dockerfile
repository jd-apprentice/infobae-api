# DEV
FROM golang:alpine3.20 AS build

WORKDIR /app
COPY . /app
RUN go mod download
RUN GOARCH=amd64 go build -o ./bin/InfobaeAPI ./src/main.go

# PROD
FROM alpine:3.20

RUN addgroup -S appgroup && adduser -S infobae -G appgroup
WORKDIR /app
COPY --from=build /app/bin/InfobaeAPI /app/bin/InfobaeAPI
ENV GIN_MODE=release

USER infobae
ENV PORT=3000
EXPOSE ${PORT}
ENTRYPOINT ["/app/bin/InfobaeAPI"]
