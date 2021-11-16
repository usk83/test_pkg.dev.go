package testpkgdevgo

const ModuleVersion = "0.1.0"

type unexported struct {
}

func New() *unexported {
	return &unexported{}
}

func Version() string {
	return ModuleVersion
}

This_line_causes_an_error
