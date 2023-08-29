package config_swagger

import (
	"fmt"
	"gopkg.in/yaml.v3"
	_ "gopkg.in/yaml.v3"
	"os"
	"warp_swagger/utils"
)

type SwaggerMaps struct {
	DefinitionsMap, ResponsesMap, PathsMap, ParametersMap map[string]any
}
type SwaggerCfg struct {
	SwagBP      *SwaggerMaps
	Swagger     string       `yaml:"swagger"`
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
	PathNode []*yaml.Node
}
type Parameters struct {
	ParametersNode []*yaml.Node
}

type Responses struct {
	ResponsesNode []*yaml.Node
}

type Definitions struct {
	DefinitionsNode []*yaml.Node
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

func NewName(prefix, toAdd string) string {
	return fmt.Sprintf("%s_%s", prefix, toAdd)
}

func (swg *SwaggerCfg) GetPaths() yaml.Node {
	node := yaml.Node{}
	return node
}

func NodeCollector(f []byte) (*yaml.Node, error) {
	var Node yaml.Node
	err := yaml.Unmarshal(f, &Node)
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}
	return &Node, nil
}

func newSwaggerMaps(definitions, responses, paths, parameters map[string]any) *SwaggerMaps {
	return &SwaggerMaps{
		DefinitionsMap: definitions,
		ResponsesMap:   responses,
		PathsMap:       paths,
		ParametersMap:  parameters,
	}
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
	CollectedNode, err := NodeCollector(f)
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	// ***   paths  *** //
	swag.Paths.PathNode = assign(CollectedNode, swag.Paths.PathNode, "paths")
	pathsMap := utils.Unwrap(swag.Paths.PathNode)

	// ***   definitions *** //
	swag.Definitions.DefinitionsNode = assign(CollectedNode, swag.Definitions.DefinitionsNode, "definitions")
	definitionsMap := utils.Unwrap(swag.Definitions.DefinitionsNode)

	// ***   responses   *** //
	swag.Responses.ResponsesNode = assign(CollectedNode, swag.Responses.ResponsesNode, "responses")
	responsesMap := utils.Unwrap(swag.Responses.ResponsesNode)

	// ***   parameters   *** //
	swag.Parameters.ParametersNode = assign(CollectedNode, swag.Parameters.ParametersNode, "parameters")
	parametersMap := utils.Unwrap(swag.Parameters.ParametersNode)

	// merging maps ...

	swag.SwagBP = newSwaggerMaps(definitionsMap, responsesMap, pathsMap, parametersMap)

	return swag, nil

}

func assign(node *yaml.Node, appendNode []*yaml.Node, assignee string) []*yaml.Node {
	for i := range node.Content[0].Content {
		if node.Content[0].Content[i].Value == assignee {
			appendNode = append(appendNode, node.Content[0].Content[i+1].Content...)
		}
	}
	return appendNode
}
