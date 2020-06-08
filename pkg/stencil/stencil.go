package stencil

import (
	"text/template"

	"github.com/konzertmeister/konzertmeister/pkg/stencil/helpers"
)

type Stencil struct {
	name      string
	arguments []string
	path      string
	template  *template.Template
}

func NewStencil(path string) (st *Stencil, err error) {
	t, err := template.ParseFiles(path)
	if err != nil {
		return nil, err
	}

	st = &Stencil{
		name:     t.Name(),
		path:     path,
		template: t,
	}

	st.parseArguments()

	return
}

func (s *Stencil) parseArguments() {
	s.arguments = helpers.WalkParseTree(s.template.Tree.Root, nil)
}

func (s *Stencil) Arguments() []string {
	return s.arguments
}
