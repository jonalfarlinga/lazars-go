package maps

import (
	"lazars-go/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

var wallSprite = assets.WallSprite

type GameMap struct {
	walls  [][]int
	sprite *ebiten.Image
}

func NewMap(xTiles, yTiles int) *GameMap {
	wallArr := Map1
	wSizeX := wallSprite.Bounds().Dx()
	wSizeY := wallSprite.Bounds().Dy()
	sprite := ebiten.NewImage(xTiles*wSizeX, yTiles*wSizeY)
	for y := 0; y < yTiles; y++ {
		for x := 0; x < xTiles; x++ {
			if wallArr[y][x] == 1 {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*wSizeX), float64(y*wSizeY))
				sprite.DrawImage(wallSprite, op)
			}
		}
	}
	return &GameMap{
		walls:  wallArr,
		sprite: sprite,
	}
}

func (m *GameMap) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(m.sprite, op)
}
