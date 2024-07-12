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
	op.GeoM.Rotate(p.position.Dir)
	op.GeoM.Translate(p.position.X-20, p.position.Y-20)
	screen.DrawImage(p.sprite, op)
}

func (p *Player) Move() {
	speed := float64(200 / ebiten.TPS())
	var Y float64 = 0
	var X float64 = 0
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		Y += speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		Y -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		X -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		X += speed
	}

	// Check for diagonal movement
	if X != 0 && Y != 0 {
		factor := speed / math.Sqrt(X*X+Y*Y)
		X *= factor
		Y *= factor
	}

	p.position.X += X
	p.position.Y += Y

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
}
