# Results API Server

## Variables

| Environment Variable | Description                | Example                                      |
|----------------------|----------------------------|----------------------------------------------|
| DB_USER              | Postgres Database user     | user                                         |
| DB_PASSWORD          | Postgres Database Password | hunter2                                      |
| DB_HOST              | Postgres Database host     | /cloudsql/my-project:us-east1:tekton-results |
| DB_NAME              | Postgres Database name     | tekton_results                               |
| DB_SSLMODE           | Database SSL mode          | verify-full                                  |
| GRPC_PORT            | gRPC Server Port           | 50051 (default)                              |
| LOG_TYPE             | Log storage type for tekton results.      | File                          |
| REST_PORT            | REST proxy Port            | 8080  (default)                              |
| PROMETHEUS_PORT      | Prometheus Port            | 9090  (default)                              |
| S3_BUCKET_NAME       | Unique S3 bucket name                     | s3_tekton-results             |
| S3_ENDPOINT          | S3 endpoint url. Points to aws by default.|                               |
| S3_REGION            | Region with S3 server                     |  eu-west-3                    |
| S3_ACCESS_KEY_ID     | Access key to S3 service endpoint         |                               |
| S3_SECRET_ACCESS_KEY | Secret access key to login to S3 endpoint |                               |
| TLS_HOSTNAME_OVERRIDE| Override the hostname used to serve TLS. This should not be set (or set to the empty string) in production environments.     | results.tekton.dev |
| TLS_PATH             | Path to TLS files          | /etc/tls                                     |

These values can also be set in the config file located in the `config/env/config` directory.
When deploying on Kubernetes, this file should be mounted in with the `tekton-results-config` ConfigMap.
Secret env variables `S3_ACCESS_KEY_ID` and `S3_SECRET_ACCESS_KEY` should be provided using k8s secret `tekton-results-s3-credentials` in the namespace tekton-pipelines.
tekton-results supports log storage types:
  - File(log content stored in the pvc)
  - S3

Values derived from Postgres DSN

If you use the default postgres database we provide, the `DB_HOST` can be set as `tekton-results-postgres-service.tekton-pipelines`.
