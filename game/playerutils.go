package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func getVectorFromKeys() (float64, float64, float64) {
	speed := float64(200 / ebiten.TPS())
	var y float64 = 0
	var x float64 = 0
	var dir float64 = 0
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

	// Handle diagonal movement
	if x != 0 && y != 0 {
		factor := speed / math.Sqrt(x*x+y*y)
		x *= factor
		y *= factor
	}
	return x, y, dir
}

func playerBounding(p *Player, dir float64) {
	bounds := p.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	if p.position.X < 0+halfW {
		p.position.X = 0 + halfW
	}
	if p.position.X > ScreenWidth-halfW {
		p.position.X = ScreenWidth - halfW
	}
	if p.position.Y < 0+halfH {
		p.position.Y = 0 + halfH
	}
	if p.position.Y > ScreenHeight-halfH {
		p.position.Y = ScreenHeight - halfH
	}

	p.position.Dir += dir
	if p.position.Dir > 2*math.Pi {
		p.position.Dir -= 2 * math.Pi
	}
	if p.position.Dir < 2*math.Pi {
		p.position.Dir += 2 * math.Pi
	}
}
