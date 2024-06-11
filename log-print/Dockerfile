FROM golang:1.22-alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download
RUN go mod verify

RUN go build -o log-print .

FROM alpine:3.19 AS runtime

COPY --from=builder /build/log-print /srv/log-print

CMD ["/srv/log-print"]
