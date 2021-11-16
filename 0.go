package testpkgdevgo

const ModuleVersion = "v0.1.7"

type unexported struct {
}

func New() *unexported {
	return &unexported{}
}

func Version() string {
	return ModuleVersion
}
