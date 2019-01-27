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
)

//state of the online rooms
const numRooms = 10
var rooms [numRooms]Room


// [START main]
func main() {
	//taking the approach of separating all words with hyphen, everything lower case
	http.HandleFunc("/refresh-lobby", refreshLobbyHandler)
	http.HandleFunc("/join", joinHandler)
	http.HandleFunc("/is-my-turn", isMyTurnHandler)
	http.HandleFunc("/most-recent-move", mostRecentMoveHandler)
	http.HandleFunc("/make-move", makeMoveHandler)
	http.HandleFunc("/undo", undoHandler)

	http.HandleFunc("/spectate", spectateHandler)

	http.HandleFunc("/exit", exitHandler)
	http.ListenAndServe(":8080", nil)
	//http.ListenAndServeTLS(":443", "ssl.crt", "ssl.key", nil)
}
// [END main]
