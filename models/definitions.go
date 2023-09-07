package models

type Definitions struct {
	Name     string
	DefProps []*DefProps
}
type DefProps struct {
	Name     string
	PropType any
}

func NewDefinition(n string, dm ...*DefProps) *Definitions {
	return &Definitions{
		Name:     n,
		DefProps: dm,
	}
}

//func NewDefProp(n string, propType any) *DefProps {
//
//}
