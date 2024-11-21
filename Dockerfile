<<<<<<< Updated upstream
# From docker hub
FROM node:lts-alpine

# Create app directory
WORKDIR /app

# Install app dependencies
COPY ./package*.json .

# Install dependencies
RUN yarn install

# Bundle app source
COPY . /app

# Build
RUN yarn build

# Expose the port
EXPOSE 8080

# Run the app
CMD ["node", "build/index.js"]
=======
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
>>>>>>> Stashed changes
