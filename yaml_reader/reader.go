package yaml_reader

import (
	"fmt"
	"warp_swagger/config_swagger"
)

type YamlReader struct {
	SwaggerCfg *config_swagger.SwaggerCfg
}

func NewReader(path string) (*YamlReader, error) {
	ymlr := &YamlReader{}
	var err error
	ymlr.SwaggerCfg, err = config_swagger.NewSwaggerCfg(path)
	if err != nil {
		return nil, fmt.Errorf("new swagger error: %w", err)
	}
	fmt.Println("swagger", ymlr.SwaggerCfg.Info.Version)
	return ymlr, nil
}
