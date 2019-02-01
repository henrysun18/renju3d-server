// Copyright 2016 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"
)

func writeJsonResponse(w http.ResponseWriter, raw interface{}) {
	w.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(raw)
	w.Write([]byte(response))
}

// refresh lobby state to check which rooms are occupied by whom
func refreshLobbyHandler(w http.ResponseWriter, r *http.Request) {
	// return a list of room states showing player names and inprogress
	summaries := [numRooms]RoomSummary{}
	for i := 0; i < numRooms; i++ {
		summaries[i] = rooms[i].Summary
	}
	writeJsonResponse(w, summaries)
}

// refresh room state to check for the presence of opponent
func refreshRoomHandler(w http.ResponseWriter, r *http.Request) {
	roomNumber, _ := strconv.Atoi(r.FormValue("room"))

	writeJsonResponse(w, rooms[roomNumber].Summary)
}

// assures the server that client is still online, otherwise a Timer goroutine will handle eviction
func keepAliveHandler(w http.ResponseWriter, r *http.Request) {
	roomNumber, _ := strconv.Atoi(r.FormValue("room"))
	playerNumber, _ := strconv.Atoi(r.FormValue("player-number"))

	room := &rooms[roomNumber]
	keepAlive(room, playerNumber)

	// If one player disconnects, evict both from the room allowing other players to play, then return error
	if room.Summary.P1 == "" && !time.Time.IsZero(room.TimeOfLastRequestFromBlack) ||
		room.Summary.P2 == "" && !time.Time.IsZero(room.TimeOfLastRequestFromWhite) {
		// Means someone recently got evicted
		writeJsonResponse(w, -1)
		room.ResetState()
	}
}

func keepAlive(room *Room, playerNumber int) {
	if (playerNumber == 1) {
		room.TimeOfLastRequestFromBlack = time.Now()
	} else if (playerNumber == 2) {
		room.TimeOfLastRequestFromWhite = time.Now()
	}
}

// join a room specifying room# and player name
func joinHandler(w http.ResponseWriter, r *http.Request) {
	roomNumber, _ := strconv.Atoi(r.FormValue("room"))
	name := r.FormValue("name")

	room := &rooms[roomNumber]
	playerNumber := 0
	if !room.IsFull() {
		playerNumber = room.AddPlayer(name)
		keepAlive(room, playerNumber)
	}
	writeJsonResponse(w, playerNumber)
}

// return true or false given room# and player#
func isMyTurnHandler(w http.ResponseWriter, r *http.Request) {
	roomNumber, _ := strconv.Atoi(r.FormValue("room"))
	playerNumber, _ := strconv.Atoi(r.FormValue("player-number"))
	room := rooms[roomNumber]

	isMyTurn := false
	if room.IsWhitesTurn && playerNumber == 2 {
		isMyTurn = true
	} else if !room.IsWhitesTurn && playerNumber == 1 {
		isMyTurn = true
	}
	//TODO if other player disconnects, u automatically win (also let client handle the frontend for this case)
	writeJsonResponse(w, isMyTurn)
}

// usually called ONCE immediately after isMyTurnHandler returns true to the client
func mostRecentMoveHandler(w http.ResponseWriter, r *http.Request) {
	roomNumber, _ := strconv.Atoi(r.FormValue("room"))
	room := rooms[roomNumber]

	writeJsonResponse(w, room.MovesHistory.Peek())
}

// usually called ONCE some time after both isMyTurn and mostRecentMove is handled
func makeMoveHandler(w http.ResponseWriter, r *http.Request) {
	roomNumber, _ := strconv.Atoi(r.FormValue("room"))
	playerNumber, _ := strconv.Atoi(r.FormValue("player-number"))
	X, _ := strconv.Atoi(r.FormValue("x"))
	Y, _ := strconv.Atoi(r.FormValue("y"))

	room := &rooms[roomNumber]
	if playerNumber == 2 && room.IsWhitesTurn || playerNumber == 1 && !room.IsWhitesTurn {
		room.MakeMove(X, Y, playerNumber)
		writeJsonResponse(w, Point{X, Y})
	} else {
		err := "don't try to make a move when it's not your turn!"
		writeJsonResponse(w, err)
		errors.New(err)
	}

}

//assumes that it's the player hitting this endpoint's turn
func undoHandler(w http.ResponseWriter, r *http.Request) {
	roomNumber, _ := strconv.Atoi(r.FormValue("room"))

	room := &rooms[roomNumber]
	room.UndoOneMove()
	writeJsonResponse(w, Point{-1, -1})
}

//returns board state
func spectateHandler(w http.ResponseWriter, r *http.Request) {
	roomNumber, _ := strconv.Atoi(r.FormValue("room"))

	writeJsonResponse(w, rooms[roomNumber].Board)
}

