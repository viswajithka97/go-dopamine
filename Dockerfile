# FROM golang:latest

# RUN mkdir /app

# ADD . /app

# WORKDIR /app

# RUN go build -o main .

# CMD ["/app/main"]



# Start from a base Go image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application's source code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port defined in the .env file
EXPOSE ${PORT}

# Set the entry point for the Docker image
ENTRYPOINT ["./main"]