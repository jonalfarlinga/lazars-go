package game

import (
	"lazars-go/assets"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

var playerSprite = assets.PlayerSprite

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Rotate(45.0 * math.Pi / 180.0)
	op.GeoM.Translate(150, 250)
	screen.DrawImage(playerSprite, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
