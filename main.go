package main

import (
	"os"

	"github.com/yaodavid1106/Blockchain_Projecy_DY/cli"
)

func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()
}
