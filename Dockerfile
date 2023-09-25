FROM golang:1.21-alpine AS builder
WORKDIR /go/src/app

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct

# copy the golang modules and start to download
COPY go.mod go.sum ./
RUN go mod download

# copy other files and build the executable
COPY . .
RUN go build -o /go/bin/url-shortener ./cmd

# use alpine image and copy builded file from another stage
FROM alpine:3.14
COPY --from=builder /go/bin/url-shortener .

ENTRYPOINT [ "./url-shortener" ]