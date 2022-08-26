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