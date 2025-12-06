FROM python:3.13-slim

# Recommended Python environment configuration for containers:
# - Do not write .pyc files
# - Do not buffer stdout/stderr (logs appear immediately)
ENV PYTHONDONTWRITEBYTECODE=1 \
    PYTHONUNBUFFERED=1

# Set the working directory inside the container.
WORKDIR /api

# Create a non-root user and group for running the application.
# Using a non-root user improves container security.
RUN groupadd -r userApp && useradd -r -g userApp userApp

# Copy dependency file first to leverage Docker layer caching.
COPY ./requirements.txt .

# Upgrade pip and install Python dependencies.
# --no-cache-dir avoids caching packages and keeps the image smaller.
RUN pip install --no-cache-dir --upgrade pip && \
    pip install --no-cache-dir -r requirements.txt

# Copy the rest of the application source code into the image.
COPY . .

# Change ownership of the application directory to the non-root user.
RUN chown -R userApp:userApp /api

# Switch to the non-root user for running the process.
USER userApp

# Expose the port where FastAPI will listen inside the container.
EXPOSE 80

# Default command to start the FastAPI application.
# The app listens on 0.0.0.0 so it is reachable from outside the container.
CMD ["fastapi", "run", "app/main.py", "--host", "0.0.0.0", "--port", "80"]
