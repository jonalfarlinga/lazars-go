package game

import (
	"lazars-go/config"
	"lazars-go/game/maps"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func getPlayerMove(speed float64) (float64, float64, float64) {
	x, y, dir := 0.0, 0.0, 0.0

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		y += speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		y -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		x -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		x += speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		dir -= 0.1
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		dir += 0.1
	}
	return x, y, dir
}

func screenBound(num, upperBound float64) float64 {
	if num < 0 {
		num = 0
	}
	if num > upperBound {
		num = upperBound
	}
	return num
}

func dirBound(dir float64) float64 {
	if dir > 2*math.Pi {
		dir -= 2 * math.Pi
	}
	if dir < 2*math.Pi {
		dir += 2 * math.Pi
	}
	return dir
}

func wallBound(x, y float64, p *Player, m *maps.GameMap) bool {
	ts := config.TileSize
	// Check if any of the corners are colliding with a wall
	if m.TileAt(int(p.topLeft.X+x)/ts, int(p.topLeft.Y+y)/ts) == 1 ||
		m.TileAt(int(p.topRight.X+x)/ts, int(p.topRight.Y+y)/ts) == 1 ||
		m.TileAt(int(p.botLeft.X+x)/ts, int(p.botLeft.Y+y)/ts) == 1 ||
		m.TileAt(int(p.botRight.X+x)/ts, int(p.botRight.Y+y)/ts) == 1 {
		return true
	}

	return false
}
