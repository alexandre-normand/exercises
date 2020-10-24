    package main

import (
	"github.com/alexandre-normand/bobino"
	"github.com/zchee/color"
)

func main() {
	choix := bobino.AskWithChoice("Est-que ton nom est Milo Normand?", "oui", "non", "Oui, et tu le sais.")
	if choix == "oui" {
		color.Yellow("cool!")
	} else if choix == "non" {
		color.Yellow("oh, c'est plate je pensais que tu était Milo Normand ")
	} else if choix == "Oui, et tu le sais." {
		color.Yellow("non c'est pas vrais ça c''est pas vrais")
	}
}
