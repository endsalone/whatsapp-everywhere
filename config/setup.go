package config

import (
	"os"
	"strconv"
)

func Setup() {
	WhatsappSetup()
	ServerSetup()
	BotSetup()
}

func getEnv(env string, defaultValue string) string {
	if value := os.Getenv(env); value != "" {
		return value
	}
	return defaultValue
}

func getIntEnv(env string, defaultValue int) int {
	value := getEnv(env, "")
	intvalue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return intvalue
}
