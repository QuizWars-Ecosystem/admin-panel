package config

import "time"

type Config struct {
	Local        bool          `env:"LOCAL" envDefault:"false"`
	LogLevel     string        `env:"LOG_LEVEL" envDefault:"info"`
	ServerURL    string        `env:"SERVER_URL" envDefault:"localhost:8082"`
	Port         int           `env:"PORT" envDefault:"8090"`
	StartTimeout time.Duration `env:"START_TIMEOUT" envDefault:"15s"`
	StopTimeout  time.Duration `env:"STOP_TIMEOUT" envDefault:"15s"`
	JWT          JWT           `envPrefix:"JWT_" envDefault:""`
}

type JWT struct {
	Secret         string        `env:"SECRET" envDefault:"secret"`
	AccessTimeout  time.Duration `env:"ACCESS_TIMEOUT" envDefault:"10m"`
	RefreshTimeout time.Duration `env:"REFRESH_TIMEOUT" envDefault:"10m"`
}
