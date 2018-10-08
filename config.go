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

// // DBMgrConfig -
// type DBMgrConfig struct {
// 	MapDB map[string]DBConfig `yaml:"mapdb"`
// }

// Config -
type Config struct {
	AddrBind   string     `yaml:"addrbind"`
	PathDBRoot string     `yaml:"pathdbroot"`
	ListDB     []DBConfig `yaml:"listdb"`
	// Cfg DBConfig
}

// NewConfig -
func NewConfig() Config {
	// dbmgr := DBMgrConfig{
	// LstDB: make(map[string]DBConfig),
	// }

	return Config{
		// ListDB: make([]DBConfig, 16),
		// Cfg:    DBConfig{},
		// DBMgr: dbmgr,
	}
}

// LoadConfig -
func LoadConfig(cfgfilename string) (*Config, error) {
	yamlFile, err := ioutil.ReadFile(cfgfilename)
	if err != nil {
		return nil, err
	}

	cfg := NewConfig()
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

// // SaveConfig -
// func SaveConfig(cfgfilename string, cfg *Config) error {
// 	buf, err := yaml.Marshal(&cfg)
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Print(string(buf))

// 	err = ioutil.WriteFile(cfgfilename, buf, os.ModeAppend)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
