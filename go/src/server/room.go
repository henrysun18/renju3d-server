package main

import (
	"errors"
	"time"

	"github.com/golang-collections/collections/stack"
)

type Room struct {
	Summary RoomSummary

	Board        [15][15]colour
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
	if X == -1 {
		// undo move
		if room.MovesHistory.Len() > 0 {
			room.IsWhitesTurn = !room.IsWhitesTurn
			undoneMove := room.MovesHistory.Pop().(Point)
			room.Board[undoneMove.X][undoneMove.Y] = neither
		}
	} else if X <= -2 {
		// -2 = undo request, sent by whoever has the turn right now
		// -3 = undo request accepted, sent by whoever now has the turn (on client side, their "turn" is to answer a pop-up)
		// -4 = undo request rejected, sent by whoever now has the turn (on client side, their "turn" is to answer a pop-up)
		room.IsWhitesTurn = !room.IsWhitesTurn
		room.MovesHistory.Push(Point{X, Y})
	} else {
		room.IsWhitesTurn = !room.IsWhitesTurn
		room.MovesHistory.Push(Point{X, Y})
		room.Board[X][Y] = colour(playerNumber)
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
