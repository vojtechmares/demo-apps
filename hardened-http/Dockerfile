FROM golang:1.23-alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download
RUN go mod verify

RUN go build -o hardened-http .

FROM alpine:3.22 AS runtime

RUN apk add --no-cache ca-certificates

RUN addgroup -S hardened && \
    adduser -S hardened -G hardened

COPY --from=builder --chown=hardened:hardened /build/hardened-http /srv/hardened-http

USER hardened

CMD ["/srv/hardened-http"]
