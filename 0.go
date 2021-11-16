package testpkgdevgo

const ModuleVersion = "0.0.4"

struct unexported struct{
}

func New() *unexported {
	return &unexported{}
}

func Version() string {
	return ModuleVersion
}
