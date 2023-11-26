package pkg

import (
	"github.com/g0r0d3tsky/parser/pkg/cfgparser"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestParseYAML(t *testing.T) {
	filePath := "test.yaml"

	testData := []byte(`port: 8080
						host: example.com
						timeout: 10`)
	err := os.WriteFile(filePath, testData, 0644)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			log.Fatal("remove temp file")
		}
	}(filePath)

	config, err := cfgparser.ParseYAML(filePath)
	if err != nil {
		t.Fatalf("failed to parse YAML: %v", err)
	}

	expectedConfig := cfgparser.Config{
		Port:    8080,
		Host:    "example.com",
		Timeout: 10,
	}

	assert.Equal(t, expectedConfig, config, "parsed config does not match expected config")
}
