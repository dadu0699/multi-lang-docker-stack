FROM node:24-alpine

#  Enable corepack and set up pnpm
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable

# Set the working directory
WORKDIR /api

# Copy dependency definitions
COPY package.json pnpm-lock.yaml ./

# Install ONLY production dependencies
RUN pnpm install --prod --frozen-lockfile

# Copy the application into the container
COPY . /api

# Create a non-root user for security
RUN adduser -D myuser

# Change ownership of the files to the non-root user
RUN chown -R myuser:myuser /api

# Switch to the non-root user to run the application
USER myuser

# Expose the port
EXPOSE 3000

# Command to run the application
CMD ["node", "src/app.js"]
