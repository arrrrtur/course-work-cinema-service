FROM golang:1.22.2-alpine3.19 AS builder

WORKDIR /cinema

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go install -a -ldflags='-w -s -extldflags "-static"' ./...

FROM alpine:3.19.1

RUN apk update && apk add --no-cache bash

WORKDIR /usr/local/bin

COPY tools/wait-for-it.sh wait-for-it.sh

RUN chmod 777 wait-for-it.sh

CMD ["course-work-cinema-service"]
