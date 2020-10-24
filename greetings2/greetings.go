package main

import (
	"github.com/alexandre-normand/bobino"
	"github.com/zchee/color"
)

func main() {
	nom := bobino.AskForInput("what's your name?")

	if nom == "Bob" || nom == "Bernadette" {
		color.Green("Greetings, %s!\n", nom)
		color.Red("Nice to 🥩 you")
	}
}
