package ankadb

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// DBConfig -
type DBConfig struct {
	Name   string `yaml:"name"`
	Engine string `yaml:"engine"`
	PathDB string `yaml:"pathdb"`
}

// Config -
type Config struct {
	AddrGRPC   string     `yaml:"addrgrpc"`
	PathDBRoot string     `yaml:"pathdbroot"`
	AddrHTTP   string     `yaml:"addrhttp"`
	ListDB     []DBConfig `yaml:"listdb"`
}

// NewConfig -
func NewConfig() *Config {
	return &Config{}
}

// LoadConfig -
func LoadConfig(cfgfilename string) (*Config, error) {
	yamlFile, err := ioutil.ReadFile(cfgfilename)
	if err != nil {
		return nil, err
	}

	cfg := NewConfig()
	err = yaml.Unmarshal(yamlFile, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
