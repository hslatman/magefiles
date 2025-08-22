package release

import (
	"fmt"
	"os"

	"github.com/magefile/mage/sh"
)

const (
	toolsMod = "./.tools/go.mod"
)

func Tools() error {
	_, err := os.Stat(toolsMod)
	if err == nil {
		return sh.RunV("go", "get", "-tool", fmt.Sprintf("-modfile=%s", toolsMod), "golang.org/x/exp/cmd/gorelease@latest")
	}

	return sh.RunV("go", "get", "-tool", "golang.org/x/exp/cmd/gorelease@latest")
}

func Check() error {
	_, err := os.Stat(toolsMod)
	if err == nil {
		return sh.RunV("go", "tool", fmt.Sprintf("-modfile=%s", toolsMod), "golang.org/x/exp/cmd/gorelease")
	}

	return sh.RunV("go", "tool", "golang.org/x/exp/cmd/gorelease")
}
