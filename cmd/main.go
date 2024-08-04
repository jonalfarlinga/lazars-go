package main

import (
	"lazars-go/game"

	"github.com/hajimehoshi/ebiten/v2"
)


func main() {
	g := &game.Game{
		PlayerA: *game.NewPlayer(),
		GameMap: *game.NewMap(8, 8),
	}
	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
