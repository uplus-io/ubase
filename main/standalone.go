package main

import (
	"fmt"
	"github.com/uplus-io/ubase"
	"github.com/uplus-io/ucluster"
	"gopkg.in/yaml.v2"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		panic("must give config file path")
	}

	configPath := args[1]

	fmt.Printf("uplus db ready read config file[%s]\n", configPath)

	config := loadConfig(configPath)

	ucluster.Serving(config.Cluster)
}

func loadConfig(configPath string) ubase.UBaseConfig {
	file, e := os.OpenFile(configPath, os.O_RDONLY, 0600)
	if e != nil {
		panic("open config file fail")
	}
	config := ubase.UBaseConfig{}
	yaml.NewDecoder(file).Decode(&config)
	return config
}
