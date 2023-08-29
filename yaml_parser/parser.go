package yaml_parser

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"warp_swagger/config_swagger"
)

type Parser struct {
	SwaggerCfg *config_swagger.SwaggerCfg
}

func NewParser(path string) (*Parser, error) {
	p := &Parser{}
	var err error
	p.SwaggerCfg, err = config_swagger.NewSwaggerCfg(path)
	if err != nil {
		return nil, fmt.Errorf("new swagger error: %w", err)
	}

	return p, nil
}

var ErrNilMap = errors.New("path map is nil")

func (p *Parser) CollectRESTmethods() error {
	m := p.SwaggerCfg.SwagBP.PathsMap
	//pm := m[title].([]*yaml.Node)
	//if pm == nil {
	//	return ErrNilMap
	//}
	for title := range m {
		pm := m[title].([]*yaml.Node)
		if pm == nil {
			return ErrNilMap
		}
		for i := range pm {
			switch pm[i].Value {
			case "post":
				fmt.Println("post detected")
			case "get":
				fmt.Println("get detected")

			}
		}

	}
	return nil
}
