package swagger

import (
	"fmt"
	"github.com/gateway-fm/warp_swagger/config_swagger"
	"github.com/go-openapi/inflect"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
	"github.com/go-swagger/go-swagger/cmd/swagger/commands/generate"
	"github.com/go-swagger/go-swagger/codescan"
	"github.com/misnaged/annales/logger"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

type ISwagger interface {
	Generate(genType string) error
}
type swaggerSpec struct {
	cfg    *config_swagger.SwaggerCfg
	Models *generate.Model
	Server *generate.Server
}

func NewSwagger(cfg *config_swagger.SwaggerCfg) ISwagger {
	return &swaggerSpec{cfg: cfg}
}
func (s *swaggerSpec) Generate(genType string) error {
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

func SpecParser(path string) (*spec.Swagger, error) {
	sp, err := loadSpec(path)
	if err != nil {
		return nil, fmt.Errorf("loadSpec finished with an error: %w", err)
	}
	opt := codescan.Options{
		InputSpec: sp,
	}

	swagSpec, err := codescan.Run(&opt)
	if err != nil {
		return nil, fmt.Errorf("codescan finished with an error: %w", err)
	}
	return swagSpec, nil
}

// TEST !!!
func swag() {
	logger.Log().Warn("start")
	cfg := "test.yaml"
	generatePath := "generate"
	var defs []string
	l, err := SpecParser(cfg)
	if err != nil {
		// todo: handle
	}
	_ = os.Mkdir(generatePath, 0777)
	GenMod(cfg, generatePath)
	for v := range l.Definitions {
		defs = append(defs, inflect.Capitalize(v))
		// fmt.Println(inflect.Capitalize(a.Title))
	}
	var srs [][]byte
	var astFile *ast.File
	for i := range GoFiles("") {
		srs = append(srs, BytesFromFile(GoFiles("")[i]))
	}
	fset := token.NewFileSet()
	var a []string
	for _, b := range srs {
		astFile, _ = parser.ParseFile(fset, "", b, parser.SkipObjectResolution)
		a = append(a, unwrapAst(astFile)...)
	}

}

// source:
// https://github.com/go-swagger/go-swagger/blob/1182d398c09304dcb6aeafa827b5cc28b0ff54b6/cmd/swagger/commands/generate/spec_go111.go#L65
func loadSpec(input string) (*spec.Swagger, error) {
	if fi, err := os.Stat(input); err == nil {
		if fi.IsDir() {
			return nil, fmt.Errorf("expected %q to be a file not a directory", input)
		}
		sp, err := loads.Spec(input)
		if err != nil {
			return nil, err
		}
		return sp.Spec(), nil
	}
	return nil, nil
}
