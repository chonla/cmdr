package cmdr

import (
	"encoding/json"
	"io/ioutil"
)

var readFileFn = ioutil.ReadFile

// Config is config
type Config struct {
	config *ConfigType
}

// ConfigType is config type
type ConfigType struct {
	// APIToken is api token
	APIToken string `json:"api_token"`

	// IP is bound IP address
	IP string `json:"ip"`

	// Port is listening port
	Port string `json:"port"`

	// Commands is list of command aliases
	Commands map[string]string `json:"commands"`
}

// NewConfig to create a new configuration
func NewConfig() *Config {
	return &Config{
		config: &ConfigType{},
	}
}

// LoadConfig to load config
func (c *Config) LoadConfig(config string) (*ConfigType, error) {
	raw, err := readFileFn(config)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(raw, c.config)
	return c.config, err
}
