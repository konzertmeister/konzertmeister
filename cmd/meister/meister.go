package main

import (
	//"fmt"

	"github.com/konzertmeister/meister/cmd/meister/command"
)

func main() {
	command.Logger = initLog()
	command.Execute()
}
