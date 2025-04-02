package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	GATEWAY_PORT      string `mapstructure:"GATEWAY_PORT"`
	AUTH_PORT         string `mapstructure:"AUTH_PORT"`
	ORGANIZATION_PORT string `mapstructure:"ORGANIZATION_PORT"`
}

func LoadConfig() (config Config, err error) {
	// Cari lokasi .env (berdasarkan working directory)
	envPath := filepath.Join(".", ".env")
	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		log.Fatalf("File .env tidak ditemukan di: %s", envPath)
	}

	viper.SetConfigFile(envPath)
	viper.AutomaticEnv() // Juga baca dari environment variables sistem

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)

	return config, err
}
