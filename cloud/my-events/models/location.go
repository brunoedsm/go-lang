package models

import(
	"gopkg.in/mgo.v2/bson"
)


type Location struct {
	ID        bson.ObjectId `bson:"_id"`
	Name      string
	Address   string
	Country   string
	OpenTime  int
	CloseTime int
	Halls     []Hall
}