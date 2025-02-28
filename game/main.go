package main

import (
	"go-game/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	// g := &game.Game{}
	g := game.NewGame()

	err := ebiten.RunGame((g))

	if err != nil {
		panic(err)
	}
}
