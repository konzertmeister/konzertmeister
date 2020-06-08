package helpers

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
)

func IsValidPackUrl(packUrl string) bool {
	pu, err := url.Parse(packUrl)

	switch pu.Scheme {
	case "":
		return isValidLocalFilepath(pu.Path)
	case "http", "https", "git":
		break
	default:
		return false
	}

	return err == nil && pu.Path != ""
}

func isValidLocalFilepath(path string) bool {
	if filepath.IsAbs(path) {
		localPath := filepath.Join(path, ".git")
		_, err := os.Stat(localPath)

		return err == nil
	} else {
		// TODO: Should prepend the default local staging dir
		localPath := filepath.Join(path, ".git")
		_, err := os.Stat(localPath)

		if err == nil {
			return true
		}
	}

	pu, err := url.Parse("http://" + path)
	fmt.Println(pu.Hostname())

	return err == nil && pu.Hostname() != ""
}
