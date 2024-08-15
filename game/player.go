package game

import (
	"lazars-go/assets"
	"lazars-go/game/maps"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	position *Vector
	sprite   *ebiten.Image
	speed    float64
	topLeft  *Vector
	topRight *Vector
	botLeft  *Vector
	botRight *Vector
}

var playerSprite = assets.PlayerSprite

func NewPlayer(x, y float64) *Player {
	sprite := playerSprite

	pos := &Vector{
		X:   x,
		Y:   y,
		Dir: 0,
	}

	player := &Player{
		position: pos,
		sprite:   sprite,
		speed:    float64(160 / ebiten.TPS()),
		topLeft:  &Vector{0, 0, 0},
		topRight: &Vector{0, 0, 0},
		botLeft:  &Vector{0, 0, 0},
		botRight: &Vector{0, 0, 0},
	}
	player.updateCorners(player.position.X, player.position.Y)

	return player
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Scale(0.75, 0.75)
	xOffset := -(float64(p.sprite.Bounds().Dx()) / 2) * .75
	yOffset := -(float64(p.sprite.Bounds().Dy()) / 2) * .75

	op.GeoM.Translate(xOffset, yOffset)
	op.GeoM.Rotate(p.position.Dir)
	op.GeoM.Translate(p.position.X-xOffset, p.position.Y-yOffset)
	screen.DrawImage(p.sprite, op)
}

func (p *Player) Move(m *maps.GameMap) {
	x, y, dir := getPlayerMove(p.speed)

	// set the new direction
	p.position.Dir += dirBound(dir)

	// Check for diagonal movement
	if x != 0 && y != 0 {
		factor := p.speed / math.Sqrt(x*x+y*y) // full speed divided by quadratic formula
		// multiply speed factor by each vector component
		x *= factor
		y *= factor
	}
	// Check collision separately for each axis
	if !wallBound(x, 0, p, m) {
		// Update X position if no collision
		p.position.X = screenBound(
			p.position.X+x,
			ScreenWidth-(float64(p.sprite.Bounds().Dx())),
		)
	}

	if !wallBound(0, y, p, m) {
		// Update Y position if no collision
		p.position.Y = screenBound(
			p.position.Y+y,
			ScreenWidth-(float64(p.sprite.Bounds().Dy())),
		)
	}

	// Update the player's corners after movement
	p.updateCorners(p.position.X, p.position.Y)
}

func (p *Player) updateCorners(x, y float64) {
	playerWidth := float64(p.sprite.Bounds().Dx()) * 0.75
	playerHeight := float64(p.sprite.Bounds().Dy()) * 0.75

	// Update each corner's position with offset
	p.topLeft.X = x
	p.topLeft.Y = y
	p.topRight.X = x + playerWidth
	p.topRight.Y = y
	p.botLeft.X = x
	p.botLeft.Y = y + playerHeight
	p.botRight.X = x + playerWidth
	p.botRight.Y = y + playerHeight
}
