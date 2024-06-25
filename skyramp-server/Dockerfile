# Use an official Node.js LTS (Long Term Support) image as the base image
FROM node:18-alpine as build

# Install curl
RUN apk add --no-cache curl

# Set the working directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json to the working directory
COPY package*.json ./

# Install project dependencies
RUN npm install

# Copy the source code to the working directory
COPY . .

# Build the TypeScript code
RUN npm run build


FROM node:18-alpine as production

# Set the working directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json to the working directory
COPY package*.json ./

# Install production dependencies
RUN npm install --production --no-audit

# Copy the built code from the previous stage
COPY --from=build /app/dist ./

# Install necessary packages (wget is already included in Alpine)
RUN apk update && apk upgrade && \
    apk add --no-cache bash github-cli

# #copy .env file
# COPY .env .

# Expose the port your application will be listening on
EXPOSE 4000

# Start the application
CMD ["node", "./index.js"]