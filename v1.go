package testpkgdevgo

struct unexported struct{
}

func New() *unexported {
	return &unexported{}
}
