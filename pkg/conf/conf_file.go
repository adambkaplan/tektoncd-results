package conf

type ConfigFile struct {
	DB_USER     string `mapstructure:"DB_USER"`
	DB_PASSWORD string `mapstructure:"DB_PASSWORD"`
	DB_PROTOCOL string `mapstructure:"DB_PROTOCOL"`
	DB_ADDR     string `mapstructure:"DB_ADDR"`
	DB_NAME     string `mapstructure:"DB_NAME"`
	DB_SSLMODE  string `mapstructure:"DB_SSLMODE"`

	LOG_TYPE string `mapstructure:"LOG_TYPE"`

	S3_BUCKET_NAME string `mapstructure:"S3_BUCKET_NAME"`

	S3_ENDPOINT          string `mapstructure:"S3_ENDPOINT"`
	S3_REGION            string `mapstructure:"S3_REGION"`
	S3_ACCESS_KEY_ID     string `mapstructure:"S3_ACCESS_KEY_ID"`
	S3_SECRET_ACCESS_KEY string `mapstructure:"S3_SECRET_ACCESS_KEY"`
}
