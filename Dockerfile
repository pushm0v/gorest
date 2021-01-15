FROM golang:1.14-alpine AS build_base

RUN apk add --no-cache git gcc libtool musl-dev ca-certificates dumb-init

WORKDIR /tmp/gorest

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -a -ldflags "-linkmode external -extldflags '-static' -s -w" -o ./out/gorest

# Start fresh from a smaller image
FROM alpine:latest
ENV TIME_ZONE=Asia/Jakarta
RUN apk add ca-certificates sqlite socat
RUN ln -snf /usr/share/zoneinfo/$TIME_ZONE /etc/localtime && echo $TIME_ZONE > /etc/timezone

COPY --from=build_base /tmp/gorest/out/gorest /app/gorest

EXPOSE 8080

CMD ["/app/gorest"]