package game

import (
	"lazars-go/assetsPkg"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}
var playerSprite = assetsPkg.PlayerSprite

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(playerSprite, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
