package stencil

import "fmt"

/*
A stencil pack is a git repository including all the stencils and directory
structure for a meister project skeleton.
*/
type StencilPack struct {
	Stencils []*Stencil

	packUrl   string
	localPath PackPath
}

func NewStencilPack(packUrl fmt.Stringer) *StencilPack {
	return &StencilPack{
		packUrl: packUrl.String(),
	}
}
