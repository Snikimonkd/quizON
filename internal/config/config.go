package config

import (
	"os"
	"quizON/internal/logger"

	"gopkg.in/yaml.v3"
)

var GlobalConfig config

const (
	configPath = "local_config.yaml"
)

func init() {
	file, err := os.Open(configPath)
	if err != nil {
		logger.Fatalf("can't open config file: %v", err)
	}
	defer func() {
		deferErr := file.Close()
		if err != nil {
			logger.Fatalf("can't close file: %v", deferErr)
		}
	}()

	d := yaml.NewDecoder(file)

	err = d.Decode(&GlobalConfig)
	if err != nil {
		logger.Fatalf("can't parse config file: %v", err)
	}
}

// Config - структура хранящая конфиг сервиса
type config struct {
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

	Cookie struct {
		// Domain - домен на котором работает сервис
		Domain string `yaml:"domain"`
		// Secure - отправлять куку или нет
		Secure bool
	} `yaml:"cookie"`
}
