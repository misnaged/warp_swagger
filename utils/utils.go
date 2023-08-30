package utils

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"warp_swagger/warp_errors"
)

type MapArrayType int

const (
	RestMethodsMAT MapArrayType = iota
	DefinitionsMAT
	ResponsesMAT
)

func Unwrap(node []*yaml.Node) map[string]any {
	nodeMap := make(map[string]any)
	for i := 0; i < len(node)-1; i += 2 {
		if len(node[i+1].Content) <= 0 {
			nodeMap[node[i].Value] = node[i+1].Value
		} else {
			nodeMap[node[i].Value] = node[i+1].Content
		}
	}
	return nodeMap
}

// EvaluateMapArraySize is needed to evaluate the number
// of maps would be created for specific node
func evaluateMapArraySize(m map[string]any, mat MapArrayType) (int, error) {
	var count int
	for title := range m {
		pm := m[title].([]*yaml.Node)
		if pm == nil {
			return 0, warp_errors.ErrNilMap
		}
		switch mat {
		case RestMethodsMAT:
			count = restMapSizeCount(pm)
		case DefinitionsMAT:
			count = len(m)
		}
	}

	return count, nil
}

func restMapSizeCount(node []*yaml.Node) int {
	var count int
	for i := range node {
		switch node[i].Value {
		case "post":
			count += 1
		case "get":
			count += 1
		}
	}
	return count
}

func PrepareMapArray(m map[string]any, mat MapArrayType) ([]map[string]any, error) {
	count, err := evaluateMapArraySize(m, mat)
	if err != nil {
		return nil, fmt.Errorf("error while counting map size: %w", err)
	}
	mapArray := make([]map[string]any, count)

	mapArray = append(mapArray[count:]) // cut empty maps
	return mapArray, nil
}
