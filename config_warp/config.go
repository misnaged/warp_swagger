package config_warp

type WarpCfg struct {
	LibraryAPI *LibraryAPI
	GatewayAPI *GatewayAPI
}

type GatewayAPI struct {
	Definitions *Definitions
}

type LibraryAPI struct {
	Definitions *Definitions
}

type Definitions struct {
	Properties *Properties
}

type Properties struct {
	PropName string
	PropType any
}
