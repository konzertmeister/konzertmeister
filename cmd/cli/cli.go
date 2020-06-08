package main

import (
	"github.com/konzertmeister/konzertmeister/cmd/cli/command"
)

func init() {
	command.SetLogger(initLog())
}

func main() {
	command.Execute()
}
