FROM golang:1.22-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download
RUN go mod verify

RUN go build -o simple-http main.go

FROM alpine:3.19 AS runtime

COPY --from=builder /build/simple-http /srv/simple-http

CMD ["/srv/simple-http"]
