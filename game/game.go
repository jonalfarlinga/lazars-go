package game

import (
	"fmt"
	"lazars-go/assetsPkg"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	fpsLastCalculated time.Time
	fpsCounter        int
	fpsText           string
}

var playerSprite = assetsPkg.PlayerSprite

func (g *Game) Update() error {
	g.fpsCounter++
	if time.Since(g.fpsLastCalculated) >= time.Second {
		g.fpsText = fmt.Sprintf("FPS: %d", g.fpsCounter)
		g.fpsCounter = 0
		g.fpsLastCalculated = time.Now()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(playerSprite, nil)
	ebitenutil.DebugPrint(screen, g.fpsText)
    // ebitenutil.DrawRect(screen, 10, 10, 100, 20, color.Black) // Background
    // text.Draw(screen, g.fpsText, mplusNormalFont, 15, 25, color.White)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
