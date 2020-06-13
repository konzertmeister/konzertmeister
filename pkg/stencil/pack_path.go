package stencil

type PackPath string

func (l PackPath) String() string {
	return string(l)
}
