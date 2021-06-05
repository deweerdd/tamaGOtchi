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
	IdleLoop    []int
	Image       *ebiten.Image
}

func NewSprite(asset string, frames int, frameWidth int, frameHeight int, idleLoop []int) Sprite {
	path := "assets/" + asset
	imgName, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatal(err)
	}

	sprite := Sprite{
		FrameWidth:  frameWidth,
		FrameHeight: frameHeight,
		FrameNum:    frames,
		Image:       imgName,
		IdleLoop:    idleLoop,
	}
	return sprite
}
