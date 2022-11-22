package cfg

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

type Config struct {
	TelegramToken string   `json:"telegram_token"`
	Database      Database `json:"database"`
}

type Database struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
}

var cfg *Config

func Get() *Config {
	if cfg == nil {
		var err error
		cfg, err = load()
		if err != nil {
			panic(err)
		}
	}

	return cfg
}

const cfgPath = "config.json"

func load() (*Config, error) {
	f, err := os.Open(cfgPath)
	if err != nil {
		return nil, err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			logrus.Errorf("cannot close config file: %v", err)
		}
	}(f)

	buf, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("cannot read config file: %w", err)
	}

	var cfg Config
	err = json.Unmarshal(buf, &cfg)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal config file: %w", err)
	}

	return &cfg, nil
}
