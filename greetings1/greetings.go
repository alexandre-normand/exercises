package main

import (
	"github.com/alexandre-normand/bobino"
	"github.com/zchee/color"
)

func main() {
	nom := bobino.AskForInput("what's your name?")
	color.Green("Greetings, %s!\n", nom)
}
