# Start from golang v1.16 base image
FROM golang:1.16-alpine as builder

RUN apk update && apk add openssh git make && rm -rf /var/cache/apk/*

WORKDIR /app
COPY Makefile .

COPY go.mod go.sum ./
RUN make deps

COPY . .
RUN make build

FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /app

COPY --from=builder /app/build/server .
COPY files/ /app/files


EXPOSE 8080
CMD ["./server"]
