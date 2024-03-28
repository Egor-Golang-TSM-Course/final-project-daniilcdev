package cfg

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	HashingServiceAddr string
	GatewayPort        string
}

func Load() *EnvConfig {
	godotenv.Load()
	cfg := EnvConfig{}
	cfg.HashingServiceAddr = fromEnvVar("HASHING_ADDRESS", ":9000")
	cfg.GatewayPort = fromEnvVar("GATEWAY_PORT", ":8080")

	return &cfg
}

func fromEnvVar(envVar string, default_ string) string {
	if val, exists := os.LookupEnv(envVar); exists {
		return val
	}

	return default_
}
