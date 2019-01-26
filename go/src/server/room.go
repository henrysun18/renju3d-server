package main

import (
	"errors"
	"github.com/golang-collections/collections/stack"
)

type Room struct {
	Summary RoomSummary

	MovesHistory stack.Stack
	IsWhitesTurn bool
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

func (room *Room) MakeMove(X int, Y int) {
	// move or undo always results in opponent's turn, except if empty board
	if (room.MovesHistory.Len() > 0) {
		room.IsWhitesTurn = !room.IsWhitesTurn
	}
	if (X == -1 && Y == -1) {
		//undo
		room.MovesHistory.Pop()
	} else {
		move := &Point{X, Y}
		room.MovesHistory.Push(move)
	}

}





type RoomSummary struct {
	P1 string //treat P1 as black for now
	P2 string
	IsInProgress bool
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