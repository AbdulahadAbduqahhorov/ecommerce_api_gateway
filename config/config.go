package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	App         string
	Environment string // dev, test, prod
	Version     string
	LogLevel string
	ServiceHost string
	HTTPPort    string

	DefaultOffset string
	DefaultLimit  string
	
	ProductServiceGrpcHost string
	ProductServiceGrpcPort string

	CategoryServiceGrpcHost string
	CategoryServiceGrpcPort string

	OrderServiceGrpcHost string
	OrderServiceGrpcPort string
}

// Load ...
func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.App = cast.ToString(getOrReturnDefaultValue("PROJECT_NAME", "Article"))
	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", "dev"))
	config.Version = cast.ToString(getOrReturnDefaultValue("VERSION", "1.0.0"))
	config.LogLevel = cast.ToString(getOrReturnDefaultValue("LOG_LEVEL", "debug"))


	config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":8080"))

	config.DefaultOffset = cast.ToString(getOrReturnDefaultValue("DEFAULT_OFFSET", "0"))
	config.DefaultLimit = cast.ToString(getOrReturnDefaultValue("DEFAULT_LIMIT", "10"))

	config.ProductServiceGrpcHost = cast.ToString(getOrReturnDefaultValue("PRODUCT_SERVICE_GRPC_HOST", "localhost"))
	config.ProductServiceGrpcPort = cast.ToString(getOrReturnDefaultValue("PRODUCT_SERVICE_GRPC_PORT", ":9001"))

	config.CategoryServiceGrpcHost = cast.ToString(getOrReturnDefaultValue("CATEGORY_SERVICE_GRPC_HOST", "localhost"))
	config.CategoryServiceGrpcPort = cast.ToString(getOrReturnDefaultValue("CATEGORY_SERVICE_GRPC_PORT", ":9001"))

	config.OrderServiceGrpcHost = cast.ToString(getOrReturnDefaultValue("ORDER_SERVICE_GRPC_HOST", "localhost"))
	config.OrderServiceGrpcPort = cast.ToString(getOrReturnDefaultValue("ORDER_SERVICE_GRPC_PORT", ":9002"))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
