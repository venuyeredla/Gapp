# syntax=docker/dockerfile:1
FROM golang:1.23.0 AS build
# FROM ubuntu:latest
#FROM alpine:3.14

# Set destination for COPY
WORKDIR /app

# Copy the source code. Note the slash at the end, as explained in. Keep slash at the end.
COPY . ./

# Download Go modules
RUN go mod download

#COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/GoApp


#FROM scratch
FROM alpine:3.14

WORKDIR /app
RUN mkdir -p static
RUN mkdir -p logs
COPY ./wstatic/*   /app/wstatic/
COPY --from=build /app/GoApp /app/GoApp

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
EXPOSE 2024

# Run
CMD ["./GoApp"]