



FROM golang:alpine AS builder
WORKDIR /build

COPY go.mod .
RUN go mod download
RUN go build -o crm.shopdev.com ./cmd/server

FROM scratch
COPY ./config /config
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/crm.shopdev.com /

ENTRYPOINT ["/crm.shopdev.com", "config/local.yaml"]


