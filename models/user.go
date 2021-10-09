package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID       bson.ObjectId `json:"id" bson:"_id"`
	Name     string        `json:"name" bson:"name"`
	Email    string        `json:"email" bson:"email"`
	Password string        `json:"password" bson:"password"`
}
type Post struct {
	ID        bson.ObjectId `json:"pid" bson:"_pid"`
	Caption   string        `json:"caption" bson:"caption"`
	ImageURL  string        `json:"iurl" bson:"iurl"`
	TimeStamp time.Time     `json:"timestamp" bson:"timestamp"`
}
