FROM golang:alpine3.20 AS build

WORKDIR /app
COPY src go.mod go.sum /app/
COPY docs /app/docs
RUN go mod download && \
    go build -o ./bin/InfobaeAPI ./src/main.go

FROM alpine:3.20

RUN addgroup -S appgroup && adduser -S infobae -G appgroup
WORKDIR /app
COPY --from=build /app/bin/InfobaeAPI /app/InfobaeAPI
COPY --from=build /app/docs /app/docs

USER infobae
ENV GIN_MODE=release
EXPOSE 3000
ENTRYPOINT ["/app/InfobaeAPI"]
