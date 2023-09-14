package swagger

import (
	"strings"
)

type GenerationType int

const (
	ServerGen GenerationType = iota
	ModelsGen
	ClientGen // TODO
	undefined
)

var genTypes = [...]string{
	ServerGen: "server",
	ModelsGen: "models",
	ClientGen: "client",
}

func (s GenerationType) String() string {
	return genTypes[s]
}

func ParseGenType(s string) GenerationType {
	for i, r := range genTypes {
		if strings.ToLower(s) == r {
			return GenerationType(i)
		}
	}
	return undefined
}
