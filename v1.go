package testpkgdevgo

const ModuleVersion = "0.0.5"

type unexported struct {
}

func New() *unexported {
	return &unexported{}
}

func Version() string {
	return ModuleVersion
}
