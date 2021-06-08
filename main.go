package main

import (
	_ "image/png"
	"strconv"
	"time"

	"github.com/ddeweerd/tamaGOtchi/components"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	count       int
	randIdle    int
	currentTime int
}

var gd = newGameData()

func (g *Game) Update() error {
	g.count++
	g.randIdle = RandIdle()
	g.currentTime = time.Now().Second()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	egg := components.NewSprite("egg.png",
		2,
		16,
		16,
		[]int{0, 1})
	i := g.count % egg.FrameNum
	drawFrame(screen, egg, egg.IdleLoop[i], 0, 9, 0)

	baby := components.NewSprite("baby.png",
		8,
		16,
		16,
		[]int{0, 1, 2, 7, 2, 1, 0, 1})
	ib := g.count % baby.FrameNum
	drawFrame(screen, baby, baby.IdleLoop[ib], 0, g.randIdle, 0)

	//ebitenutil.DebugPrint(screen, strconv.Itoa(i))
	ebitenutil.DebugPrint(screen, strconv.Itoa(g.currentTime))
}

func (g *Game) Layout(w, h int) (int, int) {
	return gd.ScreenWidth, gd.ScreenHeight
}

func main() {
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowSize(gd.ScreenWidth*2, gd.ScreenHeight*2)
	ebiten.SetMaxTPS(1)
	ebiten.SetWindowTitle("TamaGOtchi")
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
