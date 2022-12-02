package main

import (
	"log"
	"os"
)

func main() {
	// открываем файл с конфигом
	file, err := os.Open("configPath")
	if err != nil {
		log.Fatalf("can't open file: %v", err)
	}
	defer file.Close()
}
