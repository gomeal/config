package config

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v2"
)

var _ ConfigClient = &configClientImpl{}

type configClientImpl struct {
	values map[Key]Value
}

func newConfigClient(configPath string) *configClientImpl {
	c := &configClientImpl{
		values: make(map[Key]Value),
	}

	wd, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("unable to os.Getwd() in config.NewProvider: %w", err))
	}

	cfgPath, err := filepath.Abs(filepath.Join(wd, configPath))
	if err != nil {
		panic(fmt.Errorf("unable to filepath.Abs() in config.NewProvider: %w", err))
	}

	data, err := os.ReadFile(cfgPath)
	if err != nil {
		panic(fmt.Errorf("unable to os.ReadFile() in config.NewProvider, filePath: %v, err: %w", cfgPath, err))
	}

	type rawEntry struct {
		Type  string      `yaml:"type"`
		Value interface{} `yaml:"value"`
	}

	var raw map[string]rawEntry
	if err := yaml.Unmarshal(data, &raw); err != nil {
		panic(fmt.Errorf("unable to yaml.Unmarshal() in config.NewProvider: %w", err))
	}

	for key, entry := range raw {
		switch entry.Type {
		case "int":
			v, ok := entry.Value.(int)
			if !ok {
				panic(fmt.Errorf("unable to convert %v variable to type int", key))
			}

			c.saveValue(Key(key), newValue(v))
		case "string":
			v, ok := entry.Value.(string)
			if !ok {
				panic(fmt.Errorf("unable to convert %v variable to type string", key))
			}

			c.saveValue(Key(key), newValue(v))
		case "bool":
			v, ok := entry.Value.(bool)
			if !ok {
				panic(fmt.Errorf("unable to convert %v variable to type bool", key))
			}

			c.saveValue(Key(key), newValue(v))
		case "duration":
			s, ok := entry.Value.(string)
			if !ok {
				panic(fmt.Errorf("unable to convert %v variable to type string", key))
			}

			d, err := time.ParseDuration(s)
			if err != nil {
				panic(fmt.Errorf("unable to parse %v variable value %v to type duration", key, s))
			}

			c.saveValue(Key(key), newValue(d))
		default:
			panic(fmt.Errorf("unsupported type %v for key %v", entry.Type, key))
		}
	}

	return c
}

func (c *configClientImpl) GetValue(key Key) Value {
	v, ok := c.values[key]
	if !ok {
		panic(fmt.Sprintf("unable to find %v config value", key))
	}

	return v
}

func (c *configClientImpl) saveValue(key Key, v Value) {
	c.values[key] = v
}
