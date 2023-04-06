package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`

	DriveAuth struct {
		CredentialsFilePath string `yaml:"credentials_file_path"`
		TokenPath           string `yaml:"path_to_save_token"`
	} `yaml:"drive_auth"`
}

func NewConfig() *Config {
	return &Config{}
}

func (cfg *Config) ReadConfigFile() (*Config, error) {
	f, err := os.Open("../../config.yaml")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
