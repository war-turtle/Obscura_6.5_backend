package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User type to store user info
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string
	Email    string
	College  string
	Phone    string
	Onboard  bool
	TeamID   primitive.ObjectID
}

type Requests struct {
	ID   primitive.ObjectID
	Name string
}

type Team struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string
	Level     int
	CreatorID primitive.ObjectID
	Requests  []Requests
}

type Level struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Number   int32
	Name     string
	Ans      []string
	UrlAlias string
	Js       string
	Html     string
	Final    bool
}
