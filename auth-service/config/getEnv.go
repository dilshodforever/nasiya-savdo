package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HTTPPort     string
	GrpcUserPort string

	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string

	RedisDB       int
	RedisHost     string
	RedisPassword string
	RedisPort     string

	SmtpSender   string
	SmtpPassword string

	DefaultOffset string
	DefaultLimit  string

	TokenKey string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":8080"))
	config.GrpcUserPort = cast.ToString(getOrReturnDefaultValue("GRPC_USER_PORT", ":8082"))

	config.SmtpSender = cast.ToString(getOrReturnDefaultValue("SMTP_SENDER", "example@gmail.com"))
	config.SmtpPassword = cast.ToString(getOrReturnDefaultValue("SMTP_PASSWORD", "1234"))

	config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localhost"))
	config.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
	config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "postgres"))
	config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "1234"))
	config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "ll_user"))

	config.RedisHost = cast.ToString(getOrReturnDefaultValue("REDIS_HOST", "localhost"))
	config.RedisPort = cast.ToString(getOrReturnDefaultValue("REDIS_PORT", ":6379"))
	config.RedisDB = cast.ToInt(getOrReturnDefaultValue("REDIS_DB", 0))
	config.RedisPassword = cast.ToString(getOrReturnDefaultValue("REDIS_PASSWORD", ""))

	config.DefaultOffset = cast.ToString(getOrReturnDefaultValue("DEFAULT_OFFSET", "0"))
	config.DefaultLimit = cast.ToString(getOrReturnDefaultValue("DEFAULT_LIMIT", "10"))
	config.TokenKey = cast.ToString(getOrReturnDefaultValue("TokenKey", "my_secret_key"))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
