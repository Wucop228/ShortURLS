package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Config_DB struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBSSLMode  string
	ServerPort string
}

type Config_REDIS struct {
	RedisAddr     string
	RedisPassword string
	RedisDB       int
}

type Config struct {
	Config_REDIS
	Config_DB
	TTL time.Duration
}

type CacheConfig struct {
	TTL time.Duration `yaml:"ttl"`
}

type AppConfig struct {
	Cache CacheConfig `yaml:"cashe"`
}

func LoadYaml(path string) (*AppConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg AppConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load .env: %w", err)
	}

	cfg, err := LoadYaml("configs/main.yml")
	if err != nil {
		return nil, fmt.Errorf("Failed to load config: %v", err)
	}

	return &Config{
		Config_REDIS: Config_REDIS{
			RedisAddr:     os.Getenv("REDIS_ADDRESS"),
			RedisPassword: os.Getenv("REDIS_PASSWORD"),
			RedisDB: func() int {
				db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
				return db
			}(),
		},
		Config_DB: Config_DB{
			DBUser:     os.Getenv("DB_USER"),
			DBPassword: os.Getenv("DB_PASSWORD"),
			DBName:     os.Getenv("DB_NAME"),
			DBPort:     os.Getenv("DB_PORT"),
			DBSSLMode:  os.Getenv("DB_SSLMODE"),
			ServerPort: os.Getenv("SERVER_PORT"),
		},
		TTL: cfg.Cache.TTL,
	}, nil
}
