ARG GO_VERSION=1.15-alpine3.12
ARG FROM_IMAGE=alpine:3.11

FROM golang:${GO_VERSION} AS builder

RUN apk update && \
  apk add ca-certificates gettext git make && \
  rm -rf /tmp/* && \
  rm -rf /var/cache/apk/* && \
  rm -rf /var/tmp/*

COPY ./ /app

WORKDIR /app

RUN make build-for-container

FROM ${FROM_IMAGE}

RUN apk update && \
  apk add ca-certificates gettext jq curl openssl git postgresql && \
  rm -rf /tmp/* && \
  rm -rf /var/cache/apk/* && \
  rm -rf /var/tmp/*

COPY --from=builder /app/dist/db-backup /bin/

ENTRYPOINT ["db-backup"]
