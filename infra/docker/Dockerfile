FROM golang:alpine3.20 AS build

WORKDIR /app
COPY . /app
RUN go mod download && \
    go install github.com/swaggo/swag/cmd/swag@latest && \
    go build -o ./bin/InfobaeAPI ./src/main.go && \
    swag init -d src

FROM alpine:3.20

RUN addgroup -S appgroup && adduser -S infobae -G appgroup
WORKDIR /app
COPY --from=build /app/bin/InfobaeAPI /app/bin/InfobaeAPI
COPY --from=build /app/docs /app/docs
ENV GIN_MODE=release

USER infobae
ENV PORT=3000
EXPOSE ${PORT}
ENTRYPOINT ["/app/bin/InfobaeAPI"]
