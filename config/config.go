package config

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	HTTPPort       string
	DbURL          string
	SecretKey      string
	DurationExpire time.Duration
}

func Get() (Config, error) {
	return load()
}

func load() (Config, error) {
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}
	return Config{
		HTTPPort:       defaultValueString("8080", viper.GetString("HTTP_PORT")),
		DbURL:          viper.GetString("DB_URL"),
		SecretKey:      viper.GetString("JWT_SECRET"),
		DurationExpire: defaultValueTimeDurationFromString(time.Minute*15, viper.GetString("DURATION_EXPIRE")),
	}, nil
}

func defaultValueTimeDurationFromString(defaultValue time.Duration, data string) time.Duration {
	if data == "" {
		return defaultValue
	}

	res, err := time.ParseDuration(data)
	if err != nil {
		return defaultValue
	}

	return res
}

func defaultValueString(defaultValue, data string) string {
	if data == "" {
		return defaultValue
	}
	return data
}
