FROM golang:alpine3.23 AS build

WORKDIR /app
COPY src /app/src
COPY go.mod go.sum /app/
COPY docs /app/docs
RUN go mod download
RUN go mod tidy

WORKDIR /app/src
RUN go build -o /app/bin/InfobaeAPI main.go

FROM alpine:3.23

RUN addgroup -S appgroup && adduser -S infobae -G appgroup
WORKDIR /app
COPY --from=build /app/bin/InfobaeAPI /app/InfobaeAPI
COPY --from=build /app/docs /app/docs

USER infobae
ENV GIN_MODE=release
EXPOSE 3000
ENTRYPOINT ["/app/InfobaeAPI"]
