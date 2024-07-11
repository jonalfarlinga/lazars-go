package main

import (
	"lazars-go/assetsPkg"
	"lazars-go/game"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game = game.Game
var assets = assetsPkg.EmbeddedAssets

func main() {
	g := &Game{}
	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
