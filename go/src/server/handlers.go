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

// join a room specifying room# and player name
func joinHandler(w http.ResponseWriter, r *http.Request) {
	roomNumber, _ := strconv.Atoi(r.FormValue("room"))
	name := r.FormValue("name")
	if name == "" {
		errors.New("pls don't join with empty name")
	}
	room := &rooms[roomNumber] //set room as pointer to our global room so room.AddPlayer mutates it

	playerNumber := 0
	if !room.IsFull() {
		playerNumber = room.AddPlayer(name)
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
	writeJsonResponse(w, isMyTurn)
}

// usually called ONCE immediately after isMyTurnHandler returns true to the client
func mostRecentMoveHandler(w http.ResponseWriter, r *http.Request) {
	roomNumber, _ := strconv.Atoi(r.FormValue("room"))
	room := rooms[roomNumber]

	writeJsonResponse(w, room.MovesHistory.Peek())
}

// usually called ONCE some time after both isMyTurn and mostRecentMove is handled
// assumes that it's the player hitting this endpoint's turn
func makeMoveHandler(w http.ResponseWriter, r *http.Request) {
	roomNumber, _ := strconv.Atoi(r.FormValue("room"))
	X, _ := strconv.Atoi(r.FormValue("x"))
	Y, _ := strconv.Atoi(r.FormValue("y"))

	room := &rooms[roomNumber]
	room.MakeMove(X, Y)
	writeJsonResponse(w, Point{X, Y})
}

//assumes that it's the player hitting this endpoint's turn
func undoHandler(w http.ResponseWriter, r *http.Request) {
	roomNumber, _ := strconv.Atoi(r.FormValue("room"))

	room := &rooms[roomNumber]
	room.MakeMove(-1, -1)
	writeJsonResponse(w, Point{-1, -1})
}

