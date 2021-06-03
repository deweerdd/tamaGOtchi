package main

import (
	"fmt"
	"image"
	_ "image/png"
	"strconv"

	"github.com/ddeweerd/tamaGOtchi/components"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	count    int
	randIdle int
}

func drawFrame(screen *ebiten.Image, sprite components.Sprite, frameX int, frameY int, posX int, posY int) {
	gd := newGameData()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(posX*16), float64(gd.ScreenHeight)/2)
	fmt.Println(posX*16, float64(gd.ScreenWidth)/2)
	//Need to refactor for larger Sprites or Multirow Sprites.
	spriteFromSheet := sprite.Image.SubImage(image.Rect(frameX*sprite.FrameWidth, frameY, frameX*sprite.FrameWidth+sprite.FrameWidth, sprite.FrameHeight))
	screen.DrawImage(spriteFromSheet.(*ebiten.Image), op)
}

func (g *Game) Update() error {
	g.count++
	g.randIdle = RandIdle()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	sprite := components.NewSprite("baby.png", 8)
	i := g.count % sprite.FrameNum
	imageIndex := []int{0, 1, 2, 7, 2, 1, 0, 1}
	//positionIndex := []int{8, 10, 12, 14, 12, 10, 8, 10}
	//position := RandIdle()
	debugCount := strconv.Itoa(i)
	ebitenutil.DebugPrint(screen, debugCount)
	drawFrame(screen, sprite, imageIndex[i], 0, g.randIdle, 0)
}

func (g *Game) Layout(w, h int) (int, int) {
	gd := newGameData()
	return gd.ScreenWidth, gd.ScreenHeight
}

func main() {
	gd := newGameData()
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowSize(gd.ScreenWidth*2, gd.ScreenHeight*2)
	ebiten.SetMaxTPS(1)
	ebiten.SetWindowTitle("TamaGOtchi")
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
