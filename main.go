package main

import (
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

func (g *Game) Update() error {
	g.count++
	g.randIdle = RandIdle()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	sprite := components.NewSprite("baby.png",
		8,
		16,
		16,
		[]int{0, 1, 2, 7, 2, 1, 0, 1})
	i := g.count % sprite.FrameNum

	debugCount := strconv.Itoa(i)
	ebitenutil.DebugPrint(screen, debugCount)
	drawFrame(screen, sprite, sprite.IdleLoop[i], 0, g.randIdle, 0)
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
