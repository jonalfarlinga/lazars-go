package game

import (
	"lazars-go/assets"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	position Vector
	sprite   *ebiten.Image
}

var playerSprite = assets.PlayerSprite

func NewPlayer() *Player {
	sprite := playerSprite

	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

  	pos := Vector{
		X: ScreenWidth/2 - halfW,
		Y: ScreenHeight/2 - halfH,
		Dir: 1,
	}

	return &Player{
		position: pos,
		sprite: sprite,
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(
		-(float64(p.sprite.Bounds().Dx())/2),
		-(float64(p.sprite.Bounds().Dy())/2),
	)
	op.GeoM.Rotate(p.position.Dir)
	op.GeoM.Translate(p.position.X, p.position.Y)
	screen.DrawImage(p.sprite, op)
}

func (p *Player) Move() {
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

	// Check for diagonal movement
	if x != 0 && y != 0 {
		factor := speed / math.Sqrt(x*x+y*y)
		x *= factor
		y *= factor
	}

	p.position.X += x
	p.position.Y += y

	if p.position.X < 0 {
		p.position.X = 0
	}
	if p.position.X > ScreenWidth {
		p.position.X = ScreenWidth
	}
	if p.position.Y < 0 {
		p.position.Y = 0
	}
	if p.position.Y > ScreenHeight {
		p.position.Y = ScreenHeight
	}

	p.position.Dir += dir
	if p.position.Dir > 2*math.Pi {
		p.position.Dir -= 2*math.Pi
	}
	if p.position.Dir < 2*math.Pi {
		p.position.Dir += 2*math.Pi
	}
}
