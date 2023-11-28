FROM golang:1.18-alpine as build

WORKDIR /app
COPY . .
RUN ls -la
RUN go build -o server.bin cmd/main.go
CMD ./server.bin

#FROM alpine:3.18
#COPY --from=build app/server.bin .
#CMD ./server.bin