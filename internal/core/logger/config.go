package core_logger

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Level  string `envconfig:"LEVEL" required:"true"`
	Folder string `envconfig:"FOLDER" required:"true"`
}

func NewLoggerConfig() (Config, error) {
	var conf Config

	if err := envconfig.Process("LOGGER", &conf); err != nil {
		return Config{}, fmt.Errorf("process envconfig: %w", err)
	}

	return conf, nil
}

func NewConfigMust() Config {
	conf, err := NewLoggerConfig()
	if err != nil {
		err = fmt.Errorf("get logger config: %w", err)
		panic(err)
	}

	return conf
}
