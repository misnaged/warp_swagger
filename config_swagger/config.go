package config_swagger

import (
	"fmt"
	"gopkg.in/yaml.v3"
	_ "gopkg.in/yaml.v3"
	"os"
)

type SwaggerCfg struct {
	Swagger string `yaml:"swagger"`
	//NodesMap    map[string][]*yaml.Node
	Paths       *Paths       `yaml:"paths,flow"`
	Info        *Info        `yaml:"info,flow"`
	BasePath    string       `yaml:"basePath"`
	Schemes     []string     `yaml:"schemes"`
	Consumes    []string     `yaml:"consumes"`
	Produces    []string     `yaml:"produces"`
	Parameters  *Parameters  `yaml:"parameters,flow"`
	Responses   *Responses   `yaml:"responses,flow"`
	Definitions *Definitions `yaml:"definitions,flow"`
}
type Paths struct {
	// TODO
}

// **** only for demo *** //
// Info not needed to be unmarshalled

type Info struct {
	Version        string   `yaml:"version"`
	Title          string   `yaml:"title"`
	Description    string   `yaml:"description"`
	TermsOfService string   `yaml:"termsOfService"`
	Contact        *Contact `yaml:"contact,flow"`
	License        *License `yaml:"license,flow"`
}

type Contact struct {
	Name string `yaml:"name"`
}
type License struct {
	Name string `yaml:"name"`
}

type Schemes struct {
}

type Consumes struct {
}

type Produces struct {
}

type Parameters struct {
}

type Responses struct {
}

type Definitions struct {
}

//func (swg *SwaggerCfg) GetMap(path string) map[string][]*yaml.Node {
//	f, err := os.ReadFile(path)
//	if err != nil {
//		log.Fatalf("error: %v", err)
//	}
//	var newNode yaml.Node
//	err = yaml.Unmarshal(f, &newNode)
//	if err != nil {
//		log.Fatalf("error: %v", err)
//	}
//	root := newNode.Content[0].Content
//
//	nodesMap := AddToMap(root, swg.NodesMap)
//
//	return nodesMap
//}
//func AddToMap(root []*yaml.Node, nodesMap map[string][]*yaml.Node) map[string][]*yaml.Node {
//	for i := 0; i < len(root)-1; i += 2 {
//		nodesMap[root[i].Value] = nil
//	}
//	for i := 1; i < len(root)-1; i += 2 {
//		nodesMap[root[i-1].Value] = root[i].Content
//	}
//	return nodesMap
//}

func NewName(prefix, toAdd string) string {
	return fmt.Sprintf("%s_%s", prefix, toAdd)
}

func (swg *SwaggerCfg) GetPaths() yaml.Node {
	node := yaml.Node{}
	return node
}

func NewSwaggerCfg(path string) (*SwaggerCfg, error) {
	var swag *SwaggerCfg

	f, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}
	err = yaml.Unmarshal(f, &swag)
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}
	fmt.Println("swagger", swag.Swagger)

	//return &SwaggerCfg{
	//	NodesMap: make(map[string][]*yaml.Node),
	//}
	return swag, nil

}
