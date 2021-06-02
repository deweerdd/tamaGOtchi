package components

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sprite struct {
	FrameWidth  int
	FrameHeight int
	FrameNum    int
	Image       *ebiten.Image
}

func NewSprite(asset string, frames int) Sprite {
	path := "assets/" + asset
	imgName, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatal(err)
	}

	sprite := Sprite{
		FrameWidth:  16,
		FrameHeight: 16,
		FrameNum:    frames,
		Image:       imgName,
	}
	return sprite
}
