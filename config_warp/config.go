package config_warp //nolint:all

type Warp struct {
	External *ExternalPkg
	Handlers *ExternalPkg
}
type ExternalPkg struct {
	Models      []*Models
	Output      string
	PackageName string
	ProtoName   string
	ProtoPath   string
	PackageURL  string
}

type Models struct {
	Name string
	Type string //   optional in the most cases
}
