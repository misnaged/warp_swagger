package utils

import "gopkg.in/yaml.v3"

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
