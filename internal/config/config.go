package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	URL      string `mapstructure:"url"`
	Token    string `mapstructure:"token"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

func Load() (*Config, error) {
	v := viper.New()

	configDir, err := os.UserConfigDir()
	if err != nil {
		configDir = filepath.Join(os.Getenv("HOME"), ".config")
	}

	v.SetConfigName("config")
	v.SetConfigType("toml")
	v.AddConfigPath(filepath.Join(configDir, "vikunja-cli"))

	v.SetEnvPrefix("")
	v.BindEnv("url", "VIKUNJA_URL")
	v.BindEnv("token", "VIKUNJA_TOKEN")
	v.BindEnv("username", "VIKUNJA_USERNAME")
	v.BindEnv("password", "VIKUNJA_PASSWORD")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("reading config: %w", err)
		}
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("parsing config: %w", err)
	}

	return &cfg, nil
}
