package main

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	PostgreSQLDBURL string `yaml:"postgresql"`
}

func getConfig() (Config, error) {
	config := Config{}
    
	{
		var data []byte;
		var err error;

		{
			configPathEnv := os.Getenv("CONFIG_PATH")

			if len(configPathEnv) == 0 {
				return Config{}, errors.New("environment variable \"CONFIG_PATH\" should not be empty");
			} 

			data, err = os.ReadFile(configPathEnv)
			if err != nil {
				return Config{}, err
			}
		}

		err = yaml.Unmarshal(data, &config)
		if err != nil {
			return Config{}, err
		}
	}

	return config, nil
}