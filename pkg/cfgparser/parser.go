package cfgparser

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Config struct {
	Port    string `yaml:"port"`
	Host    string `yaml:"host"`
	Timeout int    `yaml:"timeout"`
}

func ParseYAML(filePath string) (Config, error) {
	var config Config

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "port:") {
			var port string
			_, err := fmt.Sscanf(line, "port: %d", &port)
			if err != nil {
				return config, fmt.Errorf("failed to parse port: %w", err)
			}
			config.Port = port
		} else if strings.HasPrefix(line, "host:") {
			var host string
			_, err := fmt.Sscanf(line, "host: %s", &host)
			if err != nil {
				return config, fmt.Errorf("failed to parse port: %w", err)
			}
			config.Host = host
		} else if strings.HasPrefix(line, "timeout:") {
			var timeout int
			_, err := fmt.Sscanf(line, "timeout: %d", &timeout)
			if err != nil {
				return config, fmt.Errorf("failed to parse timeout: %w", err)
			}
			config.Timeout = timeout
		}
	}

	return config, nil
}
