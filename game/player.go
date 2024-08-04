package game

import (
	"lazars-go/assets"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	position Vector
	sprite   *ebiten.Image
	speed	 float64
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
		speed: float64(200 / ebiten.TPS()),
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	xOffset := -(float64(p.sprite.Bounds().Dx())/2)
	yOffset := -(float64(p.sprite.Bounds().Dy())/2)

	op.GeoM.Translate(xOffset, yOffset)
	op.GeoM.Rotate(p.position.Dir)
	op.GeoM.Translate(p.position.X - xOffset, p.position.Y - yOffset)
	screen.DrawImage(p.sprite, op)
}

func (p *Player) Move() {
	x, y, dir := getMove(p.speed)

	// Check for diagonal movement
	if x != 0 && y != 0 {
		factor := p.speed / math.Sqrt(x*x+y*y)
		x *= factor
		y *= factor
	}

	// set the new X,Y position
	p.position.X = screenBound(
		p.position.X + x,
		ScreenWidth - (float64(p.sprite.Bounds().Dx())),
	)
	p.position.Y = screenBound(
		p.position.Y + y,
		ScreenWidth - (float64(p.sprite.Bounds().Dy())),
	)

	// set the new direction
	p.position.Dir += dirBound(dir)
}

func getMove(speed float64) (float64, float64, float64) {
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
		dir -= 2*math.Pi
	}
	if dir < 2*math.Pi {
		dir += 2*math.Pi
	}
	return dir
}
