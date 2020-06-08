package stencil

/*
A stencil pack is a tar archive including all the stencils and directory
structure for a meister project skeleton.
*/
type StencilPack struct {
	Stencils []*Stencil

	packUrl string
}

func NewStencilPack(packUrl string) *StencilPack {
	return &StencilPack{
		packUrl: packUrl,
	}
}
