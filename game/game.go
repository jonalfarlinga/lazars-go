package game

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	fpsLastCalculated time.Time
	fpsCounter        int
	fpsText           string
	PlayerA            Player
}

const (
	ScreenHeight = 800
	ScreenWidth = 600
)


func (g *Game) Update() error {
	g.PlayerA.Move()
	g.fpsCounter++
	if time.Since(g.fpsLastCalculated) >= time.Second {
		g.fpsText = fmt.Sprintf("FPS: %d", g.fpsCounter)
		g.fpsCounter = 0
		g.fpsLastCalculated = time.Now()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.PlayerA.Draw(screen)
	ebitenutil.DebugPrint(screen, g.fpsText)
	// ebitenutil.DrawRect(screen, 10, 10, 100, 20, color.Black) // Background
	// text.Draw(screen, g.fpsText, mplusNormalFont, 15, 25, color.White)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
