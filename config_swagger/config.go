package config_swagger

import "github.com/go-openapi/spec"

type SwaggerCfg struct {
	Spec *spec.Swagger
}

func (s *SwaggerCfg) GetSpec() {
	s.Spec.Definitions
}
