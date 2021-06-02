package main

type GameData struct {
	ScreenWidth  int
	ScreenHeight int
}

func newGameData() GameData {
	gd := GameData{
		ScreenWidth:  320,
		ScreenHeight: 240,
	}
	return gd
}
