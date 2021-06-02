package main

import (
	"image"
	_ "image/png"
	"log"
	"strconv"

	"github.com/ddeweerd/tamaGOtchi/components"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	count int
}

func drawFrame(img *ebiten.Image, sprite components.Sprite, frameX int, frameY int, canvasX int, canvasY int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(canvasX), float64(canvasY))
	img.DrawImage(sprite.Image.SubImage(image.Rect(frameX*sprite.FrameWidth, frameY*sprite.FrameHeight, sprite.FrameWidth, sprite.FrameHeight)).(*ebiten.Image), op)
}

func (g *Game) Update() error {
	g.count++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	baby, _, err := ebitenutil.NewImageFromFile("assets/baby.png")
	if err != nil {
		log.Fatal(err)
	}
	sprite := components.Sprite{
		FrameOX:     0,
		FrameOY:     0,
		FrameWidth:  16,
		FrameHeight: 16,
		FrameNum:    4,
		Image:       baby,
	}
	//op := &ebiten.DrawImageOptions{}
	//gd := newGameData()
	//op.GeoM.Scale(2, 2)
	//op.GeoM.Translate(float64((gd.ScreenWidth/2)-(sprite.FrameWidth/2)), float64(gd.ScreenHeight/2))
	//println(float64((gd.ScreenWidth/2)-(sprite.FrameWidth/2)), float64(gd.ScreenHeight/2))
	i := g.count % sprite.FrameNum
	sx, sy := sprite.FrameOX+i*sprite.FrameWidth, sprite.FrameOY
	//println(i, sx, sy, sx+sprite.FrameWidth, sy+sprite.FrameHeight)
	//screen.DrawImage(sprite.Image.SubImage(image.Rect(sx, sy, sx+sprite.FrameWidth, sy+sprite.FrameHeight)).(*ebiten.Image), op)
	drawFrame(screen, sprite, sx, sy, 0, 0)
	debugCount := strconv.Itoa(i)
	ebitenutil.DebugPrint(screen, debugCount)
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
