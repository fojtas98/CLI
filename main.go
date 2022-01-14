package main

import (
	"github.com/fojtas98/dailyMenus/cmd"
	"github.com/fojtas98/dailyMenus/data"
)

func main() {
	data.OpenDatabase()
	cmd.Execute()
}
