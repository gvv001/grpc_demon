FROM golang:1.23-alpine
RUN apk update --no-cache && apk add --no-cache tzdata
WORKDIR /app
CMD mkdir config
COPY /bin/grpc_server_linux.exe /app
COPY config /app/config
EXPOSE 8080/tcp
CMD ["./grpc_server_linux.exe"]

