FROM golang:1.13-alpine
WORKDIR /goversion
COPY . .
RUN \
    go version && \
    go build -o should-fail .
