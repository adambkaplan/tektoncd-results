package conf

type ConfigFile struct {
	DB_ADDR               string `mapstructure:"DB_ADDR"`
	DB_HOST               string `mapstructure:"DB_HOST"`
	DB_NAME               string `mapstructure:"DB_NAME"`
	DB_PASSWORD           string `mapstructure:"DB_PASSWORD"`
	DB_PORT               string `mapstructure:"DB_PORT"`
	DB_PROTOCOL           string `mapstructure:"DB_PROTOCOL"`
	DB_SSLMODE            string `mapstructure:"DB_SSLMODE"`
	DB_USER               string `mapstructure:"DB_USER"`
	GRPC_PORT             string `mapstructure:"GRPC_PORT"`
	LOG_CHUNK_SIZE        int    `mapstructure:"LOG_CHUNK_SIZE"`
	LOG_TYPE              string `mapstructure:"LOG_TYPE"`
	LOGS_DATA             string `mapstructure:"LOGS_DATA"`
	PROMETHEUS_PORT       string `mapstructure:"PROMETHEUS_PORT"`
	REST_PORT             string `mapstructure:"REST_PORT"`
	S3_BUCKET_NAME        string `mapstructure:"S3_BUCKET_NAME"`
	S3_ENDPOINT           string `mapstructure:"S3_ENDPOINT"`
	S3_HOSTNAME_IMMUTABLE bool   `mapstructure:"S3_HOSTNAME_IMMUTABLE"`
	S3_REGION             string `mapstructure:"S3_REGION"`
	S3_ACCESS_KEY_ID      string `mapstructure:"S3_ACCESS_KEY_ID"`
	S3_SECRET_ACCESS_KEY  string `mapstructure:"S3_SECRET_ACCESS_KEY"`
	S3_MULTI_PART_SIZE    int64  `mapstructure:"S3_MULTI_PART_SIZE"`
	TLS_HOSTNAME_OVERRIDE string `mapstructure:"TLS_HOSTNAME_OVERRIDE"`
	TLS_PATH              string `mapstructure:"TLS_PATH"`
}
