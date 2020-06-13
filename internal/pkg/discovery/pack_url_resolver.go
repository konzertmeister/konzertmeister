package discovery

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/konzertmeister/konzertmeister/pkg/stencil"
)

type packUrlResolver interface {
	resolveRelativePackUrl(fmt.Stringer) (fmt.Stringer, error)
	resolvePackUrlInDir(string, fmt.Stringer) (fmt.Stringer, error)
}

type packUrlResolverContext struct{}

func (r packUrlResolverContext) resolveRelativePackUrl(parsedUrl fmt.Stringer) (fmt.Stringer, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, &ResolveLocalPackUrlError{packUrl: parsedUrl.String()}
	}

	return r.resolvePackUrlInDir(currentDir, parsedUrl)
}

func (r packUrlResolverContext) resolvePackUrlInDir(parentDir string, parsedUrl fmt.Stringer) (fmt.Stringer, error) {
	resolvedPath := filepath.Join(parentDir, parsedUrl.String())
	if _, err := os.Stat(resolvedPath); err != nil {
		return nil, &ResolveLocalPackUrlError{packUrl: parsedUrl.String()}
	}

	absolutePath, err := filepath.Abs(parsedUrl.String())
	if err != nil {
		return nil, &ResolveLocalPackUrlError{packUrl: parsedUrl.String()}
	}

	return stencil.PackPath(absolutePath), nil
}
