package main

import (
	"github.com/misnaged/annales/logger"
	"warp_swagger/yaml_reader"
)

func main() {
	path := "get.yaml"
	_, err := yaml_reader.NewReader(path)
	if err != nil {
		logger.Log().Errorf("%v", err)
	}
}
