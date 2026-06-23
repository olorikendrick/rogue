package cell

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	config, err := LoadConfig("./testData/config.json")
	if err != nil {
		t.Fatalf("Failed to load Config %v", err)
	}

	if config.ID != "agent-1" {
		t.Fatalf("expected id : 'agent-1', got id: '%s'", config.ID)
	}
	if len(config.Deps) != 2 {
		t.Fatalf("Expected 2 dependencies, got: %d dependencies", len(config.Deps))
	}
	if config.Deps[0].Name != "python" {
		t.Fatalf("expected first dep 'python', got '%s'", config.Deps[0].Name)
	}

}

func TestValidate(t *testing.T) {
	config, err := LoadConfig("./testData/config.json")
	if err != nil {
		t.Fatalf("Failed to load Config %v", err)
	}
	config.Timeout = 0
	err = config.Validate()
	if err == nil {
		t.Fatalf("Expectes config validation to fail when config timeout <0")
	}
}
