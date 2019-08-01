/*
 * Copyright (c) 2019 uplus.io
 */

package main

import (
	"fmt"
	"github.com/uplus-io/ubase"
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

	file, e := os.OpenFile(configPath, os.O_RDONLY, 0600)
	if e != nil {
		panic("open config file fail")
	}
	config := &ubase.UBaseConfig{}
	yaml.NewDecoder(file).Decode(config)

	base := ubase.NewBase(config)
	base.Serving()
}
