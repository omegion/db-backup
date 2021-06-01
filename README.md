# Database Backup Tool

[![GithubBuild](https://img.shields.io/github/workflow/status/omegion/db-backup/Code%20Check)](http://pkg.go.dev/github.com/omegion/db-backup)
[![Coverage Status](https://coveralls.io/repos/github/omegion/db-backup/badge.svg?branch=master)](https://coveralls.io/github/omegion/db-backup?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/omegion/db-backup)](https://goreportcard.com/report/github.com/omegion/db-backup)
[![GoDoc](https://img.shields.io/badge/pkg.go.dev-doc-blue)](http://pkg.go.dev/github.com/omegion/db-backup)

Dump or Import backups from local or S3 buckets.

## Install Globally

```shell
go get github.com/omegion/db-backup
```

```shell
CLI command to create backup for database in local or cloud.

Usage:
  db-backup [command]

Available Commands:
  help        Help about any command
  local       dump Management
  s3          dump Management
  version     Print the version/build number

Flags:
  -h, --help              help for db-backup
      --logLevel string   Set the logging level. One of: debug|info|warn|error (default "info")

Use "db-backup [command] --help" for more information about a command.
```

## Custom S3 Endpoint for Scaleway

```shell
export AWS_ACCESS_KEY_ID=X
export AWS_SECRET_ACCESS_KEY=X
export AWS_DEFAULT_REGION=fr-par
export BUCKET_NAME=test

db-backup dump s3 \
  --type=postgres \
  --host=example.com \
  --port=1234 \
  --database=test \
  --username=test \
  --password="12345" \
  --bucket-name=$BUCKET_NAME \
  --endpoint-url=s3.fr-par.scw.cloud
```