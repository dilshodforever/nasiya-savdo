package config

import (
  "fmt"
  "os"

  "github.com/joho/godotenv"
  "github.com/spf13/cast"
)

type Config struct {
  HTTPPort string

  PostgresHost     string
  PostgresPort     int
  PostgresUser     string
  PostgresPassword string
  PostgresDatabase string

  DefaultOffset string
  DefaultLimit  string

  TokenKey string
}


func Load() Config {
  if err := godotenv.Load(); err != nil {
    fmt.Println("No .env file found")
  }

  config := Config{}

  config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":8087"))

  config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localhost"))
  config.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
  config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "postgres"))
  config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "1234"))
  config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "game"))

  config.DefaultOffset = cast.ToString(getOrReturnDefaultValue("DEFAULT_OFFSET", "0"))
  config.DefaultLimit = cast.ToString(getOrReturnDefaultValue("DEFAULT_LIMIT", "10"))
  config.TokenKey=cast.ToString(getOrReturnDefaultValue("TokenKey", "my_secret_key"))
  return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
  val, exists := os.LookupEnv(key)

  if exists {
    return val
  }

  return defaultValue
}
