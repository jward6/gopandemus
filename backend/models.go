package backend

import (
	"google.golang.org/appengine/datastore"
	"time"
)

type dbModel struct {
	key       *datastore.Key `json:"-" datastore:"-"`
	parentKey *datastore.Key `json:"-" datastore:"-"`
	created   time.Time
}

func (m *dbModel) Key() *datastore.Key {
	return m.key
}

type sessionModel struct {
	dbModel
	ended time.Time
}

type boardStateModel struct {
	dbModel
	lastModified time.Time
}

type playerModel struct {
	dbModel
	Name     string
	Location string
	Hand     []cardModel
}

type cardModel struct {
	dbModel
	Name  string
	Color string
}
