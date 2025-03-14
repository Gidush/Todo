package config

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
}

func Load() (Config, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return Config{}, errors.New("can not find path to config.go")
	}
	configDir := filepath.Dir(filename)

	confFilePath := filepath.Join(configDir, "../../build/env/local.yml")
	confFile, err := os.Open(confFilePath)
	if err != nil {
		return Config{}, err
	}
	defer confFile.Close()

	var config Config
	err = yaml.NewDecoder(confFile).Decode(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

type Server struct {
	Port string `yaml:"port"`
}

type Database struct {
	ConnString string `yaml:"connString"`
}
