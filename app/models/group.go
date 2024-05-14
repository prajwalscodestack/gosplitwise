package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Group struct {
	Id        primitive.ObjectID `json:"id" bson:"id"`
	GroupName string             `json:"groupName" bson:"groupName"`
	Members   []Member           `json:"members" bson:"members"`
}

type Member struct {
	Id     primitive.ObjectID `json:"id" bson:"id"`
	Mobile string             `json:"mobile" bson:"mobile"`
}
