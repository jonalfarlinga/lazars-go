package main

import (
	"lazars-go/game"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game = game.Game

func main() {
	g := &Game{}
	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
