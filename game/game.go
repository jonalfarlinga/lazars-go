package game

import (
	"lazars-go/game/maps"
	"lazars-go/config"
	"lazars-go/game/player"
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	fpsLastCalculated time.Time
	fpsCounter        int
	fpsText           string
	PlayerA           player.Player
	GameMap           maps.GameMap
}

const (
	ScreenHeight = config.ScreenHeight
	ScreenWidth = config.ScreenWidth
)

func NewGame(m [][]int) *Game {
	gameMap := *maps.NewMap(m)
	player := *player.NewPlayer(float64(gameMap.PlayerAStart[0]), float64(gameMap.PlayerAStart[1]))
	return &Game{
		GameMap: gameMap,
		PlayerA: player,
	}
}

func (g *Game) Update() error {
	g.PlayerA.Move(&g.GameMap)
	g.fpsCounter++
	if time.Since(g.fpsLastCalculated) >= time.Second {
		g.fpsText = fmt.Sprintf("FPS: %d", g.fpsCounter)
		g.fpsCounter = 0
		g.fpsLastCalculated = time.Now()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.GameMap.Draw(screen)
	g.PlayerA.Draw(screen)
	ebitenutil.DebugPrint(screen, g.fpsText)

	// ebitenutil.DrawRect(screen, 10, 10, 100, 20, color.Black) // Background
	// text.Draw(screen, g.fpsText, mplusNormalFont, 15, 25, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
