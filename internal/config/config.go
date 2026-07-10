package config

import "time"

type Config struct {
	App      AppConfig
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	Logger   LoggerConfig
}

type AppConfig struct {
	Name    string
	Version string
	Env     string
}

type ServerConfig struct {
	Host string
	Port int

	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
	ShutdownTimeout time.Duration
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  string
	Timezone string
}

type JWTConfig struct {
	Secret     string
	ExpireHour int
}

type LoggerConfig struct {
	Level string
}