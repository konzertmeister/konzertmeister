package helpers

import (
	"testing"
)

func TestIsValidPackUrl_valid_urls(t *testing.T) {
	validUrls := []string{
		"http://github.com/organization/repository",
		"github.com/organization/repository",
		"github.com/organization/repository@1.2.3-tag",
		"/local/file/path",
	}

	for _, url := range validUrls {
		if !IsValidPackUrl(url) {
			t.Errorf("Expected `%v` to be valid, got invalid", url)
		}
	}
}

func TestIsValidPackUrl_invalid_urls(t *testing.T) {
	validUrls := []string{
		"alskjff#?asf//dfas",
		"file:///test/hi",
	}

	for _, url := range validUrls {
		if IsValidPackUrl(url) {
			t.Errorf("Expected `%v` to be invalid, got valid", url)
		}
	}
}
