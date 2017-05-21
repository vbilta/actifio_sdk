package sdk

import (
	"os"
	"log"
	"encoding/json"
)

var ActConfig configuration

type configuration struct {
	Name, Password, VendorKey, Appliance string
}

// Initialize AppConfig
func initConfig() {
	loadConfig()
}

// Reads config.json and decode into AppConfig
func loadConfig() {
	file, err := os.Open("config.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	decoder := json.NewDecoder(file)
	ActConfig = configuration{}
	err = decoder.Decode(&ActConfig)
	if err != nil {
		log.Fatalf("[loadAppConfig]: %s\n", err)
	}
}

