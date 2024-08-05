package maps

import (
	"lazars-go/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

var wallSprite = assets.WallSprite

type GameMap struct {
	walls        [][]int
	sprite       *ebiten.Image
	PlayerAStart []int
}

func NewMap(wallArr [][]int) *GameMap {
	xTiles := len(wallArr) - 1
	yTiles := len(wallArr[1])
	wSizeX := wallSprite.Bounds().Dx()
	wSizeY := wallSprite.Bounds().Dy()
	playerAStart := wallArr[0]
	playerAStart[0] *= wSizeX
	playerAStart[1] *= wSizeY
	sprite := ebiten.NewImage(xTiles*wSizeX, yTiles*wSizeY)
	for x := 0; x < xTiles; x++ {
		for y := 0; y < yTiles; y++ {
			if wallArr[y+1][x] == 1 {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*wSizeX), float64(y*wSizeY))
				sprite.DrawImage(wallSprite, op)
			}
		}
	}
	return &GameMap{
		walls:        wallArr,
		sprite:       sprite,
		PlayerAStart: playerAStart,
	}
}

func (m *GameMap) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(m.sprite, op)
}
