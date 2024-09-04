package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	HTTPPort          string
	MongoDBConnection string
}

func Load() Config {

	config := Config{}

	config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":8089"))
	config.MongoDBConnection = cast.ToString(getOrReturnDefaultValue("MONGODB_CONNECTION", "mongodb+srv://dilshod:2514@cluster0.klxv3df.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
