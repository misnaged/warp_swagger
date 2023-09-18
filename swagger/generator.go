package swagger

import (
	"fmt"
	"github.com/gateway-fm/warp_swagger/config_swagger"
	"github.com/go-swagger/go-swagger/cmd/swagger/commands/generate"
)

type ISwagger interface {
	Generate(genType string) error
}
type swagGenerator struct {
	cfg    *config_swagger.SwaggerCfg
	Models *generate.Model
	Server *generate.Server
}

func NewSwagger(cfg *config_swagger.SwaggerCfg) ISwagger {
	return &swagGenerator{cfg: cfg}
}

func (s *swagGenerator) Generate(genType string) error {
	switch ParseGenType(genType) {
	case ClientGen:
		return ErrClientNotImplemented
	case ServerGen:
		if s.cfg.Server == nil {
			return ErrNilCfgServerOpts
		}
		server, err := GenServer(s.cfg.Server.SpecPath, s.cfg.Server.Output)
		if err != nil {
			return fmt.Errorf("failed while executing GenServer: %w", err)
		}
		s.Server = server
	case ModelsGen:
		if s.cfg.Models == nil {
			return ErrNilCfgModelsOpts
		}
		models, err := GenMod(s.cfg.Models.SpecPath, s.cfg.Models.Output)
		if err != nil {
			return fmt.Errorf("failed while executing GenMod: %w", err)
		}
		s.Models = models
	default:
		return ErrGenTypeUndefined
	}
	return nil
}
