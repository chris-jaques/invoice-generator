package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Clockify struct {
		ApiKey string `yaml:"apiKey"`
		WorkspaceId string `yaml:"workspaceId"`
		UserId string `yaml:"userId"`
	}
}

func ParseConfig() *Config{
	config := &Config{}
	file, err := os.Open("/root/invgen.conf")
	if err != nil {
		fmt.Printf("Error Opening Config:\n\t%s\n", err)
		os.Exit(1)
	}
	defer file.Close()
	
	// Init new YAML decode
	d := yaml.NewDecoder(file)
	
	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		fmt.Printf("Error parsing Config:\n\t%s\n", err)
		os.Exit(1)
	}
	return config
}
