package game

import (
	"lazars-go/assets"

	// "github.com/hajimehoshi/ebiten/v2"
)

var wallSprite = assets.WallSprite

type Map struct {
	walls	[][]int
}

func NewMap(xTiles, yTiles int) *Map {
	wallArr := make([][]int, yTiles)
	for y := range yTiles {
		wallArr[y] = make([]int, xTiles)
	}
	return &Map{
		walls: wallArr,
	}
}
