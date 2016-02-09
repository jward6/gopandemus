package backend

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"time"
)

// GameSession type
type GameSession struct {
	Key     *datastore.Key
	Created time.Time
	Ended   time.Time
}

// BoardState type
type BoardState struct {
	ID             string
	SessionID      string
	LastModified   time.Time
	Players        map[string]*Player
	InfectedCities []string
}

// Player type
type Player struct {
	Location string
	Hand     []*Card
}

// Card type
type Card struct {
	Name  string
	Color string
}

// SaveBoardState commits the state to the datastore
func SaveBoardState(c context.Context, state *BoardState) (*BoardState, error) {
	sessionKey, err := datastore.DecodeKey(state.SessionID)
	if err != nil {
		return nil, err
	}
	stateModel := &boardStateModel{}
	stateModel.parentKey = sessionKey
	stateModel.lastModified = time.Now()
	stateModel.key, err = datastore.Put(c, datastore.NewIncompleteKey(c, "BoardState", stateModel.parentKey), stateModel)
	if err != nil {
		return nil, err
	}

	// Initialize the result
	result := &BoardState{
		ID:           stateModel.key.Encode(),
		SessionID:    stateModel.parentKey.Encode(),
		LastModified: time.Now(),
		Players:      make(map[string]*Player),
	}

	// Save the players
	for k, v := range state.Players {
		p := &playerModel{
			Name:     k,
			Location: v.Location,
		}
		p.parentKey = stateModel.key
		p.key, err = datastore.Put(c, datastore.NewIncompleteKey(c, "PlayerState", p.parentKey), p)
		if err != nil {
			return nil, err
		}

		for _, card := range v.Hand {
			cardModel := &cardModel{
				Name:  card.Name,
				Color: card.Color,
			}
			cardModel.parentKey = p.key
			cardModel.key, err = datastore.Put(c, datastore.NewIncompleteKey(c, "PlayerCard", p.parentKey), cardModel)
			if err != nil {
				return nil, err
			}

		}

		// Added player to result
		result.Players[k] = &Player{
			Location: p.Location,
			Hand:     v.Hand,
		}
	}

	return result, nil
}

// NewGame creates a new BoardState
func NewGame(c context.Context) (*BoardState, error) {
	// Create and save a new game session
	session := &sessionModel{}
	session.created = time.Now()

	var err error
	session.key = datastore.NewIncompleteKey(c, "Session", nil)
	session.key, err = datastore.Put(c, session.key, session)
	if err != nil {
		return nil, err
	}

	// Initalize the BoardState for a new game
	boardState := &BoardState{
		SessionID:    session.Key().Encode(),
		LastModified: time.Now(),
		Players:      make(map[string]*Player),
	}

	// Added the Players to the BoardState
	for _, n := range []string{"Jeff", "Ben", "Dan", "Giancarlo"} {
		boardState.Players[n] = &Player{
			Location: "Atlanta",
		}

		boardState.Players[n].Hand = []*Card{&Card{
			Name:  "New York",
			Color: "Blue",
		},
		}
	}

	// Save the BoardState
	result, err := SaveBoardState(c, boardState)
	if err != nil {
		return nil, err
	}

	return result, err
}

// SaveGame saves the provided State under the provided Session
func SaveGame(c context.Context, state *BoardState) (*BoardState, error) {
	result, err := SaveBoardState(c, state)
	if err != nil {
		return nil, err
	}

	return result, nil
}
