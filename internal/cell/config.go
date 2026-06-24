package cell

import (
	"encoding/json"
	"fmt"
	"github.com/olorikendrick/rogue/internal/store"
	"os"
)

type Mount struct {
	Host      string `json:"host"`
	Container string `json:"container"`
	ReadOnly  bool   `json:"read_only"`
}

type Config struct {
	ID       string      `json:"id"`
	Priority uint        `json:"priority"`
	Commands []string    `json:"commands"`
	Network  bool        `json:"network"`
	Timeout  int         `json:"timeout"`
	Deps     []store.Dep `json:"deps"`
	Mounts   []Mount     `json:"mounts"`
}

func (c *Config) Validate() error {
	if c.Timeout <= 0 {
		return fmt.Errorf("Timeout must be greater than 0, got : %d", c.Timeout)

	}
	if c.ID == "" {
		return fmt.Errorf("Config must have a valid non empty ID")
	}
	return nil
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
