FROM python:3.13-alpine

# Set the working directory
WORKDIR /app

# Copy the requirements file and install the Python dependencies
COPY ./requirements.txt requirements.txt
RUN pip install --no-cache-dir -r requirements.txt

# Copy the application into the container
COPY . .

# Expose the port
EXPOSE 80

# Command to run the application
CMD ["fastapi", "run", "main.py", "--port", "80"]
