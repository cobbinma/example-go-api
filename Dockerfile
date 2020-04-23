# syntax = docker/dockerfile:experimental
FROM golang:1.14-stretch AS builder

ENV GO111MODULE=on
WORKDIR /src

COPY . .

RUN go test ./... \
    && CGO_ENABLED=0 go build -a -o /main cmd/api/main.go

# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
FROM alpine:3.8 as release

COPY --from=builder /main /
COPY --from=builder /src/api /api
COPY --from=builder /src/migrations /migrations

EXPOSE 8989
ENTRYPOINT ["/sbin/tini", "--"]
CMD ["/main"]
