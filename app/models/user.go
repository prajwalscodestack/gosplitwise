package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID   `json:"_id" bson:"_id"`
	Name     string               `json:"name" bson:"name" validate:"required,min=3,max=20"`
	Mobile   string               `json:"mobile" bson:"mobile" validate:"required,min=10"`
	Email    string               `json:"email" bson:"email" validate:"required,email"`
	Password string               `json:"password" bson:"password" validate:"required,required,min=5"`
	Friends  []primitive.ObjectID `json:"friends" bson:"friends"`
}

type Friendship struct {
	UserID   primitive.ObjectID `json:"userId" bson:"userId"`
	FriendID primitive.ObjectID `json:"friendId" bson:"friendId"`
}
