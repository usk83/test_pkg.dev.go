package testpkgdevgo

const ModuleVersion = "v0.1.1"

type unexported struct {
}

func New() *unexported {
	return &unexported{}
}

func Version() string {
	return ModuleVersion
}
