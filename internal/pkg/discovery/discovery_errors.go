package discovery

import "fmt"

type InvalidPackUrlError struct {
	packUrl string
}

func (e *InvalidPackUrlError) Error() string {
	return fmt.Sprintf("Invalid pack url: %v", e.packUrl)
}

type ResolvePackUrlError struct {
	packUrl string
	err     error
}

func (e *ResolvePackUrlError) Error() string {
	if e.err != nil {
		return fmt.Sprintf("Couldn't resolve pack url `%v`: %v", e.packUrl, e.err)
	}

	return fmt.Sprintf("Couldn't resolve pack url `%v`", e.packUrl)
}

type ResolveLocalPackUrlError struct {
	packUrl string
}

func (e *ResolveLocalPackUrlError) Error() string {
	return fmt.Sprintf("Couldn't resolve relative pack url: %v", e.packUrl)
}
