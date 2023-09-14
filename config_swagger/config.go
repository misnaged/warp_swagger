package config_swagger //nolint:all

type SwaggerCfg struct {
	Models *Spec
	Server *Spec
}

type Spec struct {
	Output, SpecPath string
}
