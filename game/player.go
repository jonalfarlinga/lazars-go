package game

import (
	"lazars-go/assets"

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
		X:   ScreenWidth/2 - halfW,
		Y:   ScreenHeight/2 - halfH,
		Dir: 1,
	}

	return &Player{
		position: pos,
		sprite:   sprite,
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	bounds := p.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(p.position.Dir)
	op.GeoM.Translate(halfW, halfH)
	op.GeoM.Translate(p.position.X-halfW, p.position.Y-halfH)
	screen.DrawImage(p.sprite, op)
}

func (p *Player) Move() {
	x, y, dir := getVectorFromKeys()

	p.position.X += x
	p.position.Y += y

	playerBounding(p, dir)
}
