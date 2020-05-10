package main

import "github.com/marcoshuck/go-terrain/pkg/game"

func main() {
	g := game.NewGame()

	g.Create().Run()
}
