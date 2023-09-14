package swagger

import (
	"fmt"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
	"github.com/go-swagger/go-swagger/codescan"
	"github.com/misnaged/annales/logger"
	"github.com/misnaged/warp_swagger/models"
	"os"
)

type IDummy interface {
	GetHandlersModel() *models.Handlers
	Call() error
}

type dummy struct {
	Handlers *models.Handlers
}

func NewDummy() IDummy {
	return &dummy{Handlers: models.NewHandler()}
}
func (d *dummy) Call() error {
	d.swag()
	return nil
}
func (d *dummy) GetHandlersModel() *models.Handlers {
	return d.Handlers
}

// func LogThis(path string) (*log.Logger, *os.File) {
// 	kek, err := os.Create(path)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	l := log.New(kek, "", 0)
// 	return l, kek
// }

func (d *dummy) swag() {
	logger.Log().Warn("start")
	cfg := "./example/swagger_dummy.yaml"
	l, _ := SpecParser(cfg)
	// if err != nil {
	//	todo: handle
	// }

	// dl, file1 := LogThis("defs.txt")
	// pl, file2 := LogThis("paths.txt")
	// gl, file3 := LogThis("gets.txt")
	// defer file3.Close()
	// defer file2.Close()
	// defer file1.Close()

	var paths []string

	for k := range l.Paths.Paths {
		paths = append(paths, k)
	}

	for i := range paths {
		d.Handlers.Operations = append(d.Handlers.Operations, ParseOperations(l.Paths.Paths[paths[i]]))
	}

}
func OperationCheck(oper *spec.Operation) bool {
	return oper != nil
}
func GetHandlerOutputName(operID, restMethod string) string {
	output := fmt.Sprintf("%s_%s.go", operID, restMethod)
	return output
}
func collect(collection []string, str ...string) []string {
	collection = append(collection, str...)
	return collection
}

func ParseOperations(pathItem spec.PathItem) *models.Operation {
	if OperationCheck(pathItem.Get) {

		return models.NewOperation(GetHandlerOutputName(pathItem.Get.OperationProps.ID, "get"), pathItem.Get.OperationProps.ID)
	}
	if OperationCheck(pathItem.Post) {
		return models.NewOperation(GetHandlerOutputName(pathItem.Post.OperationProps.ID, "post"), pathItem.Post.OperationProps.ID)
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

/*
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
*/
