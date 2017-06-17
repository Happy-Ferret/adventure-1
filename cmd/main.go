package main

import (
	"github.com/cosmodash/adventure/game"

	"github.com/fatih/color"
)

func main() {
	color.Cyan("Welcome to Cosmodash Adventure!\n")

	g := game.Game{}
	g.Run()
}
