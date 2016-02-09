package backend

import (
	"encoding/json"
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

func init() {
	http.HandleFunc("/api/new-game", initializeGame)
	http.HandleFunc("/api/board-state", commitBoardState)
}

func initializeGame(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	board, err := NewGame(c)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	enc := json.NewEncoder(w)
	err = enc.Encode(board)
	if err != nil {
		http.Error(w, "Failed to Create New game.", 500)
		return
	}

}

func commitBoardState(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	log.Debugf(c, "Commit Board State received!", nil)

	if r.Method != "POST" {
		http.Error(w, "Invalid method.", 500)
		return
	}

	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()

	state := &BoardState{}
	err := dec.Decode(&state)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse board state. %s", err.Error()), 500)
		return
	}
	log.Debugf(c, "Input: %v", state)

	newState, err := SaveGame(c, state)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to save board state. %s", err.Error()), 500)
		return
	}

	enc := json.NewEncoder(w)
	err = enc.Encode(newState)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse board state response. %s", err.Error()), 500)
		return
	}
}
