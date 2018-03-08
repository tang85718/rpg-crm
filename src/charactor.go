package crm

import "gopkg.in/mgo.v2/bson"

const (
	ActorDB  = "crm"
	ActorCOL = "actors"
)

type Charactor struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	PlayerToken string        `bson:"player_token"`
	Name        string        `bson:"name"`
	HP          int           `bson:"hp"`
	Energy      int           `bson:"energy"`
	EnergyType  int           `bson:"energy_type"`
}
