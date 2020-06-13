package discovery

import (
	"errors"
	"fmt"
	"net/url"
	"path/filepath"
	"testing"

	. "github.com/konzertmeister/konzertmeister/test"

	"github.com/konzertmeister/konzertmeister/pkg/stencil"
)

//----------------------------------- Mocks -----------------------------------

var (
	resolveRelativePackUrlMock func(parsedUrl fmt.Stringer) (fmt.Stringer, error)
	resolvePackUrlInDirMock    func(parentDir string, parsedUrl fmt.Stringer) (fmt.Stringer, error)
)

type packUrlResolverContextMock struct{}

func (r packUrlResolverContextMock) resolveRelativePackUrl(parsedUrl fmt.Stringer) (fmt.Stringer, error) {
	return resolveRelativePackUrlMock(parsedUrl)
}

func (r packUrlResolverContextMock) resolvePackUrlInDir(parentDir string, parsedUrl fmt.Stringer) (fmt.Stringer, error) {
	return resolvePackUrlInDirMock(parentDir, parsedUrl)
}

//------------------------------- ParsePackUrl -------------------------------

func TestParsePackUrl_absolute_path(t *testing.T) {
	expected := "/absolute/pack/path"

	parsedUrl, err := parsePackUrl(expected)
	if err != nil {
		t.Fatal(err)
	}

	_, result := parsedUrl.(stencil.PackPath)
	expected_type := true

	Eq(t, result, expected_type)
	Eq(t, parsedUrl.String(), expected)
}

func TestParsePackUrl_relative_path(t *testing.T) {
	expected := "relative.tld/path"

	parsedUrl, err := parsePackUrl(expected)
	if err != nil {
		t.Fatal(err)
	}

	_, result := parsedUrl.(stencil.PackPath)
	expected_type := true

	Eq(t, result, expected_type)
	Eq(t, parsedUrl.String(), expected)
}

func TestParsePackUrl_remote_url(t *testing.T) {
	expected := "http://hostname.tld/path"

	parsedUrl, err := parsePackUrl(expected)
	if err != nil {
		t.Fatal(err)
	}

	_, result := parsedUrl.(*url.URL)
	expected_type := true

	Eq(t, result, expected_type)
	Eq(t, parsedUrl.String(), expected)
}

//------------------------------ ResolvePackUrl ------------------------------

func TestResolvePackUrl_absolute_path(t *testing.T) {
	expected := "/absolute/pack/path"

	packPath, err := resolvePackUrl(stencil.PackPath(expected))
	if err != nil {
		t.Fatal(err)
	}

	result := packPath.String()

	Eq(t, result, expected)
}

func TestResolvePackUrl_relative_path(t *testing.T) {
	_pack_url_resolver = packUrlResolverContextMock{}

	expected := "relative.tld/path"

	resolveRelativePackUrlMock = func(parsedUrl fmt.Stringer) (fmt.Stringer, error) {
		return parsedUrl, nil
	}

	packPath, err := resolvePackUrl(stencil.PackPath(expected))
	if err != nil {
		t.Fatal(err)
	}

	result := packPath.String()

	Eq(t, result, expected)
}

func TestResolvePackUrl_relative_path_in_cache(t *testing.T) {
	_pack_url_resolver = packUrlResolverContextMock{}
	_cache_path = filepath.Join("~", ".meister", "cache")

	test_path := "relative.tld/path"
	expected := filepath.Join(_cache_path, test_path)

	resolveRelativePackUrlMock = func(parsedUrl fmt.Stringer) (fmt.Stringer, error) {
		return nil, errors.New("")
	}

	resolvePackUrlInDirMock = func(parentDir string, parsedUrl fmt.Stringer) (fmt.Stringer, error) {
		return stencil.PackPath(filepath.Join(parentDir, parsedUrl.String())), nil
	}

	packPath, err := resolvePackUrl(stencil.PackPath(test_path))
	if err != nil {
		t.Fatal(err)
	}

	result := packPath.String()

	Eq(t, result, expected)
}

func TestResolvePackUrl_relative_path_is_remote(t *testing.T) {
	_pack_url_resolver = packUrlResolverContextMock{}

	test_path := "relative.tld/path"

	resolveRelativePackUrlMock = func(parsedUrl fmt.Stringer) (fmt.Stringer, error) {
		return nil, errors.New("")
	}

	resolvePackUrlInDirMock = func(parentDir string, parsedUrl fmt.Stringer) (fmt.Stringer, error) {
		return nil, errors.New("")
	}

	packPath, err := resolvePackUrl(stencil.PackPath(test_path))
	if err != nil {
		t.Fatal(err)
	}

	packUrl, result := packPath.(*url.URL)
	expected_type := true
	expected := fmt.Sprintf("https://%v", test_path)

	Eq(t, result, expected_type)
	Eq(t, packUrl.String(), expected)
}

func TestResolvePackUrl_remote_url_in_cache(t *testing.T) {
	_pack_url_resolver = packUrlResolverContextMock{}
	_cache_path = filepath.Join("~", ".meister", "cache")

	test_path, err := url.Parse("http://hostname.tld/path")
	if err != nil {
		t.Fatal(err)
	}

	resolvePackUrlInDirMock = func(parentDir string, parsedUrl fmt.Stringer) (fmt.Stringer, error) {
		return stencil.PackPath(filepath.Join(parentDir, parsedUrl.String())), nil
	}

	packPath, err := resolvePackUrl(test_path)
	if err != nil {
		t.Fatal(err)
	}

	packUrl, result := packPath.(stencil.PackPath)
	expected_type := true
	expected := filepath.Join(_cache_path, test_path.Host, test_path.Path)

	Eq(t, result, expected_type)
	Eq(t, packUrl.String(), expected)
}

func TestResolvePackUrl_remote_url_not_in_cache(t *testing.T) {
	_pack_url_resolver = packUrlResolverContextMock{}
	_cache_path = filepath.Join("~", ".meister", "cache")

	expected := "http://hostname.tld/path"
	test_path, err := url.Parse(expected)
	if err != nil {
		t.Fatal(err)
	}

	resolvePackUrlInDirMock = func(parentDir string, parsedUrl fmt.Stringer) (fmt.Stringer, error) {
		return nil, errors.New("")
	}

	packPath, err := resolvePackUrl(test_path)
	if err != nil {
		t.Fatal(err)
	}

	packUrl, result := packPath.(*url.URL)
	expected_type := true

	Eq(t, result, expected_type)
	Eq(t, packUrl.String(), expected)
}
