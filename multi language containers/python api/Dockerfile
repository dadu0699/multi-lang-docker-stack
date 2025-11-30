FROM python:3.13-slim

# Set the working directory
WORKDIR /code

# Copy the requirements file and install the Python dependencies
COPY ./requirements.txt requirements.txt
RUN pip install --no-cache-dir -r requirements.txt

# Copy the application into the container
COPY . .

# Create a non-root user for security
RUN useradd -m myuser

# Change ownership of the files to the non-root user
RUN chown -R myuser:myuser .

# Switch to the non-root user to run the application
USER myuser

# Expose the port
EXPOSE 80

# Command to run the application
CMD ["fastapi", "run", "app/main.py", "--port", "80"]
