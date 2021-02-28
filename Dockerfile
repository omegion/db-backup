ARG GO_VERSION=1.15-alpine3.12
ARG FROM_IMAGE=alpine:3.11

FROM golang:${GO_VERSION} AS builder

ADD ./dist/db-backup-linux /app/dist/db-backup

WORKDIR /app

FROM ${FROM_IMAGE}

RUN apk update && \
  apk add ca-certificates gettext jq curl openssl git postgresql && \
  rm -rf /tmp/* && \
  rm -rf /var/cache/apk/* && \
  rm -rf /var/tmp/*

COPY --from=builder /app/dist/db-backup /bin/

ENTRYPOINT ["db-backup"]
