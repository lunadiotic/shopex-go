package config

import (
	"github.com/spf13/viper"
)

func Load() (*Config, error) {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := &Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	cfg.Server.ReadTimeout = viper.GetDuration("server.readTimeout")
	cfg.Server.WriteTimeout = viper.GetDuration("server.writeTimeout")
	cfg.Server.IdleTimeout = viper.GetDuration("server.idleTimeout")
	cfg.Server.ShutdownTimeout = viper.GetDuration("server.shutdownTimeout")

	return cfg, nil
}