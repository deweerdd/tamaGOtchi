package components

import "github.com/hajimehoshi/ebiten/v2"

type Sprite struct {
	FrameOX     int
	FrameOY     int
	FrameWidth  int
	FrameHeight int
	FrameNum    int
	Image       *ebiten.Image
}
