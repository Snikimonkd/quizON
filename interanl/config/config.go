package config

import (
	"os"
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
	config := &Config{}

	// открываем файл с конфигом
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
