package config

import (
	"fmt"
	"github.com/achillescres/pkg/utils"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"path/filepath"
)

type Config struct {
	Project struct {
		AbsPath string `env:"ABS_PATH" env-required:"true"`
	} `env-prefix:"PROJECT_"`

	Postgres struct {
		Host      string `env:"HOST" env-required:"true"`
		Port      uint   `env:"PORT" env-required:"true"`
		Username  string `env:"USERNAME" env-required:"true"`
		Password  string `env:"PASSWORD" env-required:"true"`
		DbName    string `env:"DB_NAME" env-required:"true"`
		UseCA     bool   `env:"USE_CA" env-required:"true"`
		CaPath    string `env:"CA_PATH" env-required:"true"`
		CaAbsPath string
	} `env-prefix:"POSTGRES_"`

	//S3 struct {
	//	BucketName string `env:"BUCKET_NAME" env-required:"true"`
	//} `env-prefix:"S3_"`

	//Kafka struct {
	//	BrokersRaw          string `env:"BROKERS" env-required:"true"`
	//	Brokers             []string
	//	Username            string `env:"USERNAME" env-required:"true"`
	//	Password            string `env:"PASSWORD" env-required:"true"`
	//	UseSASL             bool   `env:"USE_SASL" env-required:"true"`
	//	UseCA               bool   `env:"USE_CA" env-required:"true"`
	//	CaPath              string `env:"CA_PATH" env-required:"true"`
	//	CaAbsPath           string
	//	TimeoutMilliseconds uint `env:"TIMEOUT_MILLISECONDS" env-required:"true"`
	//	Timeout             time.Duration
	//} `env-prefix:"KAFKA_"`
}

func (c *Config) Read() error {
	ew := utils.NewErrorWrapper("Config - Read")

	err := godotenv.Load("./.env")
	if err != nil {
		return ew(fmt.Errorf("load env file: %w", err))
	}
	err = cleanenv.ReadEnv(c)
	if err != nil {
		return ew(fmt.Errorf("read env into config: %w", err))
	}

	c.Postgres.CaAbsPath = filepath.Join(c.Project.AbsPath, c.Postgres.CaPath)
	//c.Kafka.CaAbsPath = filepath.Join(c.Project.AbsPath, c.Kafka.CaPath)
	//c.Kafka.Timeout = time.Duration(c.Kafka.TimeoutMilliseconds) * time.Millisecond
	//c.Kafka.Brokers = strings.Split(c.Kafka.BrokersRaw, ",")

	return nil
}
