package main

import (
	"lazars-go/game"
	"lazars-go/game/maps"

	"github.com/hajimehoshi/ebiten/v2"
)


func main() {
	g := &game.Game{
		PlayerA: *game.NewPlayer(),
		GameMap: *maps.NewMap(8, 8),
	}
	ebiten.SetWindowSize(game.ScreenWidth*1.5, game.ScreenHeight*1.5)
	ebiten.SetWindowTitle("Lazars!")
	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
