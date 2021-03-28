package models

import(
	"gopkg.in/mgo.v2/bson"
)

type Event struct {
    ID bson.ObjectId `bson:"_id"`
    Name string
    Duration int
    StartDate int64
    EndDate int64
    Location Location
}