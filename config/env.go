package config

import (
	
	viper "github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"POSTGRES_HOST"`
	DBUser     string `mapstructure:"POSTGRES_USER"`
	DBPort     string `mapstructure:"POSTGRES_PORT"`
	DBPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName     string `mapstructure:"POSTGRES_DB"`
}

func EnvLoad(path string) (config Config, err error) {

	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	return config , nil
}
