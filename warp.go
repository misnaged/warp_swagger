package main

import (
	"fmt"
	"github.com/misnaged/annales/logger"
	"warp_swagger/yaml_parser"
)

func main() {
	path := "get.yaml"
	p, err := yaml_parser.NewParser(path)
	if err != nil {
		logger.Log().Errorf("%v", err)
	}

	m, err := p.CollectRESTmethods()
	if err != nil {
		logger.Log().Errorf("%v", err)
	}
	for i := range m {
		fmt.Println(m[i])
	}
}
