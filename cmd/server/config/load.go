package config

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"strings"
	"sync"
)

var (
	onceEnv    sync.Once
	onceConfig sync.Once
)

func Load(configFile string) (*Configuration, error) {
	var err error
	var cfg *Configuration
	onceConfig.Do(func() {
		cfg, err = loadConfiguration(configFile)
		if err != nil {
			err = fmt.Errorf("error loading configuration: %w", err)
		}
	})

	return cfg, err
}

func loadConfiguration(configFile string) (*Configuration, error) {
	var configuration Configuration
	onceEnv.Do(loadEnv)

	v := viper.New()
	v.SetConfigType("yaml")

	for k, val := range defaults {
		v.SetDefault(k, val)
	}

	// add allowed env vars
	for _, key := range allowedEnvVarKeys {
		err := viper.BindEnv(key)
		if err != nil {
			return nil, fmt.Errorf("error binding env var %s: %w", key, err)
		}
	}
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if configFile != "" {
		v.SetConfigFile(configFile)
		if err := v.ReadInConfig(); err != nil {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}

	if err := v.Unmarshal(&configuration); err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %w", err)
	}

	return &configuration, nil
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Debug("no .env file found")
	}
}
