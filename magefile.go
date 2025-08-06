//go:build mage

package magefiles

import "github.com/magefile/mage/sh"

// Tidy ensures the dependencies are available
func Tidy() error {
	return sh.RunV("go", "mod", "tidy")
}
