# First stage: Build the application
FROM maven:3.9.9-eclipse-temurin-21-alpine AS build
# Set working directory
WORKDIR /app
# Copy pom.xml and download dependencies (this helps avoid downloading dependencies if they haven’t changed)
COPY pom.xml /app/pom.xml
RUN mvn dependency:go-offline
# Copy the source code
COPY src /app/src
# Build the application (skip tests to speed up the process)
RUN mvn clean package -DskipTests


# Second stage: Create Image for runtime
FROM eclipse-temurin:21.0.1_12-jre-alpine
# Copy the JAR file from the build stage
COPY --from=build /app/target/api-0.0.1-SNAPSHOT.jar /home/api-0.0.1-SNAPSHOT.jar
# Expose the port the app will run on
EXPOSE 8080
# Command to run the application
CMD ["java", "-jar", "/home/api-0.0.1-SNAPSHOT.jar"]
