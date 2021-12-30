package main

import (
	"github.com/fojtas98/CLI/cmd"
	"github.com/fojtas98/CLI/data"
)

func main() {
	data.OpenDatabase()
	cmd.Execute()
}
