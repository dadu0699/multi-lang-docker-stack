FROM node:22-alpine

# Set the working directory
WORKDIR /api

# Copy dependency definitions
COPY package.json package-lock.json /api/

# Install dependencies
RUN npm ci --omit=dev

# Copy the application into the container
COPY . /api

# Expose the port
EXPOSE 3000

# Command to run the application
CMD ["node", "src/app.js"]
