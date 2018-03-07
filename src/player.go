package crm

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type Player struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	DisplayID  string        `bson:"id"`
	Token      string        `bson:"token"`
	Phone      string        `bson:"phone"`
	Country    string        `bson:"country"`
	Province   string        `bson:"province"`
	City       string        `bson:"city"`
	Lat        float64       `bson:"lat"`
	Lng        float64       `bson:"lng"`
	Birthday   time.Time     `bson:"birthday"`
	CreateTime time.Time     `bson:"create_time"`
	UpdateTime time.Time     `bson:"update_time"`
}
