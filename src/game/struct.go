package game

import "main/src/entity"

type menu int

const (
	HOME     menu = iota
	SETTINGS menu = iota
	PLAY     menu = iota
)

type game int

const (
	INGAME   game = iota
	GAMEOVER game = iota
)

type Game struct {
	Player	  entity.Player
	Slice     []entity.Slice
	IsRunning bool
	StateMenu menu
	StateGame game
}
