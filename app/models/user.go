package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID   `json:"id" bson:"_id"`
	Name     string               `json:"name" bson:"name"`
	Mobile   string               `json:"mobile" bson:"mobile"`
	Email    string               `json:"email" bson:"email"`
	Password string               `json:"password,omitempty" bson:"-"`
	Friends  []primitive.ObjectID `json:"friends" bson:"friends"`
}

type Friendship struct {
	UserID   primitive.ObjectID `json:"userId" bson:"userId"`
	FriendID primitive.ObjectID `json:"friendId" bson:"friendId"`
}
