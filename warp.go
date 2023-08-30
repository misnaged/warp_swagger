package main

import (
	"fmt"
	"github.com/misnaged/annales/logger"
	"gopkg.in/yaml.v3"
	"warp_swagger/yaml_parser"
)

func main() {
	path := "get.yaml"
	p, err := yaml_parser.NewParser(path)
	if err != nil {
		logger.Log().Errorf("%v", err)
	}

	m, err := p.CollectDefinitions()
	if err != nil {
		logger.Log().Errorf("%v", err)
	}

	fmt.Println(m[0]["shop"].([]*yaml.Node)[0].Value)
}
