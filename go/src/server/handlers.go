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

func writeJsonResponse(w http.ResponseWriter, response []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(response))
}

func refreshLobbyHandler(w http.ResponseWriter, r *http.Request) {
	// return a list of room states showing player names and inprogress
	summaries := [numRooms]RoomSummary{}
	for i := 0; i < numRooms; i++ {
		summaries[i] = rooms[i].Summary
	}
	response, _ := json.Marshal(summaries)
	writeJsonResponse(w, response)
}

func joinHandler(w http.ResponseWriter, r *http.Request) {
	roomNumber, _ := strconv.Atoi(r.FormValue("room"))
	name := r.FormValue("name")
	if name == "" {
		errors.New("pls don't join with empty name")
	}

	playerNumber := 0
	if !rooms[roomNumber].IsFull() {
		playerNumber = rooms[roomNumber].AddPlayer(name)
	}
	writeJsonResponse(w, []byte(strconv.Itoa(playerNumber)))
}


