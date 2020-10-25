package main

import (
	"strings"

	"github.com/alexandre-normand/bobino"
	"github.com/zchee/color"
)

func main() {
	nom := strings.ToLower(bobino.AskForInput("C'est quoi ton nom?"))
	if nom == "milo" {
		color.Yellow("cool! très bon nom!")
	}
}
