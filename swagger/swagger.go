package swagger

import (
	"fmt"
	"github.com/go-openapi/inflect"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
	"github.com/go-swagger/go-swagger/codescan"
	"github.com/misnaged/annales/logger"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func Sabaka(path string) *spec.Swagger {
	p, _ := loadSpec(path)
	opt := codescan.Options{
		InputSpec: p,
	}

	lapsha, err := codescan.Run(&opt)
	if err != nil {
		logger.Log().Errorf("error %v", err)
	}
	return lapsha
}

// TEST !!!
func swag() {
	logger.Log().Warn("start")
	cfg := "test.yaml"
	generatePath := "generate"
	var defs []string
	l := Sabaka(cfg)
	_ = os.Mkdir(generatePath, 0777)
	GenMod(cfg, generatePath)
	for v := range l.Definitions {
		//fmt.Println(inflect.Capitalize(v))
		defs = append(defs, inflect.Capitalize(v))
		//fmt.Println(inflect.Capitalize(a.Title))
		//for _, v := range zhopa.Properties {+
		//	fmt.Println(v)
		//}
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
