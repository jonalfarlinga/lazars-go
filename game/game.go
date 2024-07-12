package game

import (
	"fmt"
	"lazars-go/assets"
	"math"

	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	fpsLastCalculated time.Time
	fpsCounter        int
	fpsText           string
}

var playerSprite = assets.PlayerSprite

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
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Rotate(45.0 * math.Pi / 180.0)
	op.GeoM.Translate(150, 250)
	screen.DrawImage(playerSprite, op)
	ebitenutil.DebugPrint(screen, g.fpsText)
    // ebitenutil.DrawRect(screen, 10, 10, 100, 20, color.Black) // Background
    // text.Draw(screen, g.fpsText, mplusNormalFont, 15, 25, color.White)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
