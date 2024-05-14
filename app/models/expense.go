package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Expense struct {
	Id             primitive.ObjectID `json:"id" bson:"id"`
	GroupId        primitive.ObjectID `json:"groupId" bson:"groupId"`
	TotalAmt       int64              `json:"totalAmt" bson:"totalAmt"`
	PaidBy         string             `json:"paidBy" bson:"paidBy"`
	PayableMembers []ExpenseMeta      `json:"payableMembers" bson:"payableMembers"`
}
type ExpenseMeta struct {
	UserId primitive.ObjectID `json:"userId" bson:"userId"`
	Amount int64              `json:"amount" bson:"amount"`
	Status string             `json:"status" bson:"status"`
}
