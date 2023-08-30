package yaml_parser

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"warp_swagger/config_swagger"
	"warp_swagger/utils"
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

// countRESTmethods needed to sum all REST methods
// being provided within swagger config and return the sum number
// Basically, needed only for CollectRESTmethods
// where we need to set the length of []map[string]any
func (p *Parser) countRESTmethods() (int, error) {
	var count int
	m := p.SwaggerCfg.SwagBP.PathsMap
	for title := range m {
		pm := m[title].([]*yaml.Node)
		if pm == nil {
			return 0, ErrNilMap
		}
		for i := range pm {
			switch pm[i].Value {
			case "post":
				count += 1
			case "get":
				count += 1
			}
		}
	}
	return count, nil
}

// CollectRESTmethods searching any (post and get, at this time) kind of REST methods
// in the 'paths' section of config (where they normally should be defined)
func (p *Parser) CollectRESTmethods() ([]map[string]any, error) {
	m := p.SwaggerCfg.SwagBP.PathsMap
	count, err := p.countRESTmethods()
	if err != nil {
		return nil, fmt.Errorf("error while counting rest methods: %w", err)
	}
	restMap := make([]map[string]any, count)

	restMap = append(restMap[count:]) // cut empty maps

	for title := range m {
		pm := m[title].([]*yaml.Node)
		if pm == nil {
			return nil, ErrNilMap
		}
		for i := range pm {
			switch pm[i].Value {
			case "post":
				restMap = append(restMap, utils.Unwrap(pm))
			case "get":
				restMap = append(restMap, utils.Unwrap(pm))
			}
		}

	}
	return restMap, nil
}
