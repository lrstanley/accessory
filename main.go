package main

import (
	"os"

	"github.com/spf13/afero"

	"github.com/lrstanley/accessory/cmd"
)

func main() {
	cmd.Execute(afero.NewOsFs(), os.Args)
}
