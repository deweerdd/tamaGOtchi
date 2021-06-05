package main

import (
	"image"

	"github.com/ddeweerd/tamaGOtchi/components"
	"github.com/hajimehoshi/ebiten/v2"
)

func drawFrame(screen *ebiten.Image, sprite components.Sprite, frameX int, frameY int, posX int, posY int) {
	gd := newGameData()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(posX*16), float64(gd.ScreenHeight)/2)

	//Need to refactor for larger Sprites or Multirow Sprites.
	spriteFromSheet := sprite.Image.SubImage(image.Rect(frameX*sprite.FrameWidth, frameY, frameX*sprite.FrameWidth+sprite.FrameWidth, sprite.FrameHeight))
	screen.DrawImage(spriteFromSheet.(*ebiten.Image), op)
}
