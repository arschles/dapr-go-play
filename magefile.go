//+build mage

package main

import (
	"log"

	"github.com/magefile/mage/sh"
)

// Runs go mod download and then installs the binary.
func BuildCache() error {
	out, err := sh.Output("go", "build", "-o bin/validatecache", "./cmd/validatecache")

	log.Print(out)
	return err
}
