package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// Config - структура хранящая конфиг сервиса
type Config struct {
	// Server - конфиг сервера
	Server struct {
		// Port - порт на котором слушает сервер
		Port string `yaml:"port"`
	} `yaml:"server"`

	// Database - структурка для конфига базы
	Database struct {
		// DSN - строка для подключения к базе
		DSN string `yaml:"dsn"`
	} `yaml:"database"`
}

// NewConfig - конструктор для конфига
func NewConfig(configPath string) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer func() {
		deferErr := file.Close()
		if err != nil {
			log.Fatalf("unexpected error: %v", deferErr)
		}
	}()

	d := yaml.NewDecoder(file)

	var config Config
	err = d.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
