FROM golang:1.24-alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download
RUN go mod verify

RUN go build -o log-print .

FROM alpine:3.22 AS runtime

RUN apk add --no-cache ca-certificates

COPY --from=builder /build/log-print /srv/log-print

CMD ["/srv/log-print"]
