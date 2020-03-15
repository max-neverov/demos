FROM golang:1.14-alpine
WORKDIR /test
COPY . .
RUN \
    go version && \
    go build -o test .
