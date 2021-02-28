# Database Backup Tool

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