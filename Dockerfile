# FROM ubuntu:latest
FROM alpine:3.14

RUN mkdir -p /app/static
COPY ./wstatic/*   /app/wstatic/

WORKDIR /app

COPY ./Gapp  /app/

EXPOSE 2024

# Run
CMD ["./Gapp"]