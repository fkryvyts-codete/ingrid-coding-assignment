// Package main is the main package of the application
package main

import (
	"os"

	"github.com/fkryvyts-codete/ingrid-coding-assignment/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
