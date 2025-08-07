package targets

import (
	"fmt"
	"os"

	"github.com/magefile/mage/sh"
)

const (
	toolsMod = "./.tools/go.mod"
	toolsGo  = "tools.go"
)

// Tools ensures the tools get installed
func Tools() error {
	_, err := os.Stat(toolsMod)
	if err == nil {
		return sh.RunV("go", "mod", "tidy", fmt.Sprintf("-modfile=%s", toolsMod))
	}

	// TODO: add support for tools.go?

	return sh.RunV("go", "mod", "tidy")
}

// Tidy ensures the dependencies are available
func Tidy() error {
	return sh.RunV("go", "mod", "tidy")
}
