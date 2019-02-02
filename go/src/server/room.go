package main

import (
	"errors"
	"github.com/golang-collections/collections/stack"
	"time"
)

type Room struct {
	Summary RoomSummary

	Board [15][15]colour
	MovesHistory stack.Stack
	IsWhitesTurn bool

	TimeOfLastRequestFromBlack time.Time
	TimeOfLastRequestFromWhite time.Time
}

func (room Room) IsFull() bool {
	return room.Summary.P1 != "" && room.Summary.P2 != ""
}

// returns the player number (1 corresponds to black for now)
func (room *Room) AddPlayer(name string) int {
	if room.Summary.P1 == "" {
		room.Summary.P1 = name
		return 1
	} else if room.Summary.P2 == "" {
		room.Summary.P2 = name
		return 2
	} else {
		errors.New("don't call AddPlayer when room full pls")
		return -1
	}
}

func (room Room) HasBlackPlayerWithName(name string) bool {
	return name == room.Summary.P1
}

func (room *Room) MakeMove(X int, Y int, playerNumber int) {
	if (X == -1 && Y == -1) {
		room.UndoOneMove()
	} else {
		room.IsWhitesTurn = !room.IsWhitesTurn

		room.MovesHistory.Push(Point{X, Y})
		room.Board[X][Y] = colour(playerNumber)
	}
}

func (room *Room) UndoOneMove() {
	if (room.MovesHistory.Len() > 0) {
		room.IsWhitesTurn = !room.IsWhitesTurn
		undoneMove := room.MovesHistory.Pop().(Point)
		room.Board[undoneMove.X][undoneMove.Y] = neither
	}
}

func (room *Room) ResetState() {
	*room = Room{}
}




type RoomSummary struct {
	P1 string //treat P1 as black for now
	P2 string
}

type colour int
const (
	neither colour = iota
	black
	white
)

type Point struct {
	X int // define (-1, -1) as undo move
	Y int
}