package main

import (
	"github.com/alexandre-normand/bobino"
	"github.com/zchee/color"
)

func main() {
	n := bobino.AskForNumber("quel chiffre?")

	r := 0
	for i := 1; i <= n; i++ {
		color.Red("On ajoute %d à %d et on obtient %d", i, r, r+i)
		r = r + i
	}

	color.Green("Le total est %d\n", r)
}
