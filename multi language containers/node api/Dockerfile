FROM node:22-alpine

# Set the working directory
WORKDIR /api

# Copy the package.json and package-lock.json files, and install the dependencies
COPY package*.json /api
RUN npm install

# Copy the application into the container
COPY . /api

# Expose the port
EXPOSE 3000

# Command to run the application
CMD ["node", "src/app.js"]