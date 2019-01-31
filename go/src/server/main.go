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
	"net/http"
	"time"
)

//state of the online rooms
const numRooms = 10
var rooms [numRooms]Room


// [START main]
func main() {
	//taking the approach of separating all words with hyphen, everything lower case
	http.HandleFunc("/refresh-lobby", refreshLobbyHandler)
	http.HandleFunc("/refresh-room", refreshRoomHandler)
	http.HandleFunc("/keep-alive", keepAliveHandler)
	http.HandleFunc("/join", joinHandler)
	http.HandleFunc("/is-my-turn", isMyTurnHandler)
	http.HandleFunc("/most-recent-move", mostRecentMoveHandler)
	http.HandleFunc("/make-move", makeMoveHandler)
	http.HandleFunc("/undo", undoHandler)

	http.HandleFunc("/spectate", spectateHandler)

	go evictAbsentPlayersPeriodically()

	http.ListenAndServe(":8080", nil)
	//http.ListenAndServeTLS(":443", "ssl.crt", "ssl.key", nil)

}
// [END main]

func evictAbsentPlayersPeriodically() {
	// keepAlive is called by client every 5s; let's give keepAlive callers a chance to live until now+10s (6s should be fine too)
	//
	freshnessDuration := time.Second * 10
	waitBetweenChecksDuration := time.Second * 5 //arbitrary value less than freshnessDuration
	ticker := time.NewTicker(waitBetweenChecksDuration)
	for t := range ticker.C {
		// check if we got any requests from clients within the last 5 seconds
		// if not, that means they disconnected
		for i := range rooms {
			room := &rooms[i]
			if t.Sub(room.TimeOfLastRequestFromBlack) > freshnessDuration {
				room.Summary.P1 = ""
			}
			if t.Sub(room.TimeOfLastRequestFromWhite)  > freshnessDuration {
				room.Summary.P2 = ""
			}
		}
	}
}
