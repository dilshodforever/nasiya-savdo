package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	HTTPPort string

	TokenKey string
}

func Load() Config {

	config := Config{}

	config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":8080"))
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
