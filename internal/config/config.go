// config.go

package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"sync"
	"time"
)

type Config struct {
	IsDebug       bool `env:"IS_DEBUG" yaml:"isDebug" env-default:"false"`
	IsDevelopment bool `env:"IS_DEV" yaml:"isDevelopment" env=default:"false"`
	HTTP          struct {
		IP           string        `env:"BIND_IP" yaml:"IP" env-default:"127.0.0.2"`
		Port         int           `env:"PORT" yaml:"port" env-default:"10000"`
		ReadTimeout  time.Duration `env:"HTTP-READ-TIMEOUT" yaml:"readTimeout"`
		WriteTimeout time.Duration `env:"HTTP-WRITE-TIMEOUT" yaml:"writeTimeout"`
		CORS         struct {
			AllowedMethods     []string `env:"HTTP-CORS-ALLOWED-METHODS" yaml:"allowedMethods"`
			AllowedOrigins     []string `env:"HTTP-CORS-ALLOWED-ORIGINS" yaml:"allowedOrigins"`
			AllowCredentials   bool     `env:"HTTP-CORS-ALLOWED-CREDENTIALS" yaml:"allowCredentials"`
			AllowedHeaders     []string `env:"HTTP-CORS-ALLOWED-HEADERS" yaml:"allowedHeaders"`
			OptionsPassthrough bool     `env:"HTTP-CORS-OPTIONS-PASSTHROUGH" yaml:"optionsPassthrough"`
			ExposedHeaders     []string `env:"HTTP-CORS-EXPOSED-HEADERS" yaml:"exposedHeaders"`
			Debug              bool     `env:"HTTP-CORS-DEBUG" yaml:"debug"`
		} `yaml:"cors"`
	}
	AppConfig struct {
		LogLevel  string `env:"LOG_LEVEL" yaml:"logLevel" env-default:"trace"`
		AdminUser struct {
			Email    string `env:"ADMIN_EMAIL" yaml:"email" env-default:"admin"`
			Password string `env:"ADMIN_PWD" yaml:"password" env-default:"admin"`
		}
	}
	PostgreSQL struct {
		Password string `env:"PSQL_PASSWORD" yaml:"password"`
		Host     string `env:"PSQL_HOST" yaml:"host"`
		Port     string `env:"PSQL_PORT" yaml:"port"`
		Database string `env:"PSQL_DATABASE" yaml:"database"`
		Username string `env:"PSQL_USERNAME" yaml:"username"`
	} `yaml:"postgresql"`
}

const (
	EnvConfigPathName  = "CONFIG-PATH"
	FlagConfigPathname = "config"
)

var configPath string
var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		flag.StringVar(
			&configPath,
			FlagConfigPathname,
			"./config/config.yaml",
			"this is app config file",
		)
		flag.Parse()

		log.Print("config init")
		dir, _ := os.Getwd()
		log.Printf("Current working directory: %s", dir)

		if configPath == "" {
			configPath = os.Getenv(EnvConfigPathName)
		}

		if configPath == "" {
			log.Fatal("config path is required")
		}

		instance = &Config{}

		if err := cleanenv.ReadConfig(configPath, instance); err != nil {
			log.Printf("Error reading config: %v", err)
			helpText := "Baryspiyev Artur - Monolith Cinema System"
			help, _ := cleanenv.GetDescription(instance, &helpText)
			log.Print(help)
			log.Fatal(err)
		}
	})
	return instance
}
