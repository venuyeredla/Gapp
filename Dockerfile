# FROM ubuntu:latest
FROM alpine:3.14

RUN mkdir -p /app/static
COPY ./static/*   /app/static/

WORKDIR /app

COPY ./Gapp  /app/

EXPOSE 2023

# Run
CMD ["./Gapp"]