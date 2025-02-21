FROM golang:1.24-alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download
RUN go mod verify

RUN go build -o simple-http .

FROM alpine:3.22 AS runtime

RUN apk add --no-cache ca-certificates

COPY --from=builder /build/simple-http /srv/simple-http

CMD ["/srv/simple-http"]
