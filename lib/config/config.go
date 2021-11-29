package config

import (
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

var envConfig map[string]string

func Init() {
	envConfig, _ = godotenv.Read()
}

func GetString(key string) string {
	return envConfig[key]
}

func GetInt(key string) int {
	result, _ := strconv.Atoi(envConfig[key])
	return result
}

func GetFloat(key string) float64 {
	result, _ := strconv.ParseFloat(envConfig[key], 64)
	return result
}

func GetBool(key string) bool {
	return envConfig[key] == "true"
}
