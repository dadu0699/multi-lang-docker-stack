FROM python:3.13-alpine

# Set the working directory
WORKDIR /code

# Copy the requirements file and install the Python dependencies
COPY ./requirements.txt /code/requirements.txt
RUN pip install /code/requirements.txt

# Copy the application into the container
COPY ./app /code/app

# Expose the port
EXPOSE 80

# Command to run the application
CMD ["fastapi", "run", "app/main.py", "--port", "80"]
