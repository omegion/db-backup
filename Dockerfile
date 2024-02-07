ARG GO_VERSION=1.22-alpine3.18
ARG FROM_IMAGE=postgres:alpine3.18

FROM --platform=${BUILDPLATFORM} golang:${GO_VERSION} AS builder

ARG TARGETOS
ARG TARGETARCH
ARG VERSION

LABEL org.opencontainers.image.source="https://github.com/omegion/db-backup"

WORKDIR /app

COPY ./ /app

RUN apk update && \
  apk add ca-certificates gettext make jq curl openssl git postgresql && \
  rm -rf /tmp/* && \
  rm -rf /var/cache/apk/* && \
  rm -rf /var/tmp/*

RUN make build TARGETOS=$TARGETOS TARGETARCH=$TARGETARCH VERSION=$VERSION

FROM ${FROM_IMAGE}

COPY --from=builder /app/dist/db-backup /bin/db-backup

ENTRYPOINT ["db-backup"]
