package main

import "time"

type GameData struct {
	ScreenWidth  int
	ScreenHeight int
	startTime    int
}

func newGameData() GameData {
	gd := GameData{
		ScreenWidth:  320,
		ScreenHeight: 240,
		startTime:    time.Now().Minute(),
	}
	return gd
}
