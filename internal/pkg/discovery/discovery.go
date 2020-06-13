package discovery

import (
	"fmt"
	"net/url"
	"path/filepath"

	"github.com/konzertmeister/konzertmeister/pkg/stencil"
)

var (
	_cache_path        string
	_pack_url_resolver packUrlResolver = packUrlResolverContext{}
)

func Initialize(packUrl string, cachePath string) (*stencil.StencilPack, error) {
	_cache_path = cachePath

	parsedUrl, err := parsePackUrl(packUrl)
	if err != nil {
		return nil, err
	}

	packPath, err := resolvePackUrl(parsedUrl)
	if err != nil {
		return nil, err
	}

	return stencil.NewStencilPack(packPath), nil
}

func parsePackUrl(packUrl string) (fmt.Stringer, error) {
	pu, err := url.Parse(packUrl)
	if err != nil {
		return stencil.PackPath(""), err
	}

	switch pu.Scheme {
	case "":
		cleanPath := filepath.Clean(packUrl)
		return stencil.PackPath(cleanPath), nil
	case "http", "https", "git":
		return pu, nil
	}

	return stencil.PackPath(""), &InvalidPackUrlError{packUrl: packUrl}
}

func resolvePackUrl(parsedUrl fmt.Stringer) (fmt.Stringer, error) {
	switch parsedUrl.(type) {
	case stencil.PackPath:
		// Absolute path
		if filepath.IsAbs(parsedUrl.String()) {
			return parsedUrl, nil
		}

		// Relative path
		resolvedPath, err := _pack_url_resolver.resolveRelativePackUrl(parsedUrl)
		if err != nil {
			cachePath, err := _pack_url_resolver.resolvePackUrlInDir(_cache_path, parsedUrl)
			if err == nil {
				// In Cache
				return cachePath, nil
			}

			// Is a remote path
			remoteUrl, err := url.Parse(fmt.Sprintf("https://%v", parsedUrl.String()))
			if err != nil {
				return nil, &ResolvePackUrlError{packUrl: remoteUrl.String(), err: err}
			}

			return remoteUrl, nil
		}

		// In current PWD
		return resolvedPath, nil
	case *url.URL:
		// Check cache first
		packUrl := parsedUrl.(*url.URL)

		localPath := filepath.Join(packUrl.Host, packUrl.Path)
		cachePath, err := _pack_url_resolver.resolvePackUrlInDir(_cache_path, stencil.PackPath(localPath))
		if err != nil {
			// Not in cache, get it remotely
			return parsedUrl, nil
		}

		// Found in cache
		return cachePath, nil
	}

	return nil, &ResolvePackUrlError{packUrl: parsedUrl.String()}
}
