FROM golang:1.24-alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download
RUN go mod verify

RUN go build -o https-passthrough .

FROM alpine:3.20 AS runtime

RUN apk add --no-cache ca-certificates

COPY --from=builder /build/https-passthrough /srv/https-passthrough

CMD ["/srv/https-passthrough"]
