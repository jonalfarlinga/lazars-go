package assetsPkg

import (
	"image"
	_ "image/png"
	"github.com/hajimehoshi/ebiten/v2"
)

var assets = EmbeddedAssets
var PlayerSprite = mustLoadImage("image/player.png")

func mustLoadImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}
