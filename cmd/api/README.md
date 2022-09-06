# Results API Server

## Variables

| Environment Variable | Description                     | Example                                      |
|----------------------|-------------------------------------------|----------------------------------------------|
| DB_USER              | MySQL Database user                       | user                                         |
| DB_PASSWORD          | MySQL Database Password                   | hunter2                                      |
| DB_PROTOCOL          | MySQL Database Network Protocol           | unix                                         |
| DB_ADDR              | MySQL Database address                    | /cloudsql/my-project:us-east1:tekton-results |
| DB_NAME              | MySQL Database name                       | tekton_results                               |
| DB_SSLMODE           | Database SSL mode                         | verify-full                                  |
| LOG_TYPE             | Log storage type for tekton results.      | File                                         |
| S3_BUCKET_NAME       | Unique S3 bucket name                     | s3_tekton-results                            |
| S3_ENDPOINT          | S3 endpoint url. Points to aws by default.|                                              |
| S3_REGION            | Region with S3 server                     |  eu-west-3                                   |
| S3_ACCESS_KEY_ID     | Access key to S3 service endpoint         |                                              |
| S3_SECRET_ACCESS_KEY | Secret access key to login to S3 endpoint |                                              |

These env values can be set using config map "tekton-results-config" in the namespace tekton-pipelines.
Secret env variables S3_ACCESS_KEY_ID and S3_SECRET_ACCESS_KEY should be provided using k8s secret "tekton-results-s3-credentials" in the namespace tekton-pipelines.
tekton-results supports log storage types:
  - File(log content stored in the pvc)
  - S3

Notice: For development purose you can execute controllers with outside cluster configuration. To set up env variables you can use "env/config" file.

Values derived from MySQL DSN (see
https://github.com/go-sql-driver/mysql#dsn-data-source-name)

If you use the default MySQL server we provide, the `DB_ADDR` can be set as `tekton-results-mysql.tekton-pipelines`.
