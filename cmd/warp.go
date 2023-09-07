package main

import (
	"fmt"
	"github.com/misnaged/annales/logger"
	"warp_swagger/yaml_parser"
)

func main() {
	path := "../get.yaml"
	p, err := yaml_parser.NewParser(path)
	if err != nil {
		logger.Log().Errorf("%v", err)
	}

	//m, err := p.NewDefinition()
	//if err != nil {
	//	logger.Log().Errorf("%v", err)
	//}
	m := p.NewDefinition()
	fmt.Println(m.Name)
	//fmt.Println(m[0]["shop"].([]*yaml.Node)[1].Value)

}
