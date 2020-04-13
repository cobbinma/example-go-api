# syntax = docker/dockerfile:experimental
FROM golang:1.14-stretch AS builder

ENV GO111MODULE=on
WORKDIR /src

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod,sharing=locked \
    go test ./... \
    && CGO_ENABLED=0 go build -a -o /main cmd/api/main.go

# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
FROM alpine:3.8 as release

RUN --mount=type=cache,target=/var/cache/apk apk add --update \
  curl \
  tini \
  ;

COPY --from=builder /main /
COPY --from=builder /src/files /files

EXPOSE 8989
ENTRYPOINT ["/sbin/tini", "--"]
CMD ["/main"]
