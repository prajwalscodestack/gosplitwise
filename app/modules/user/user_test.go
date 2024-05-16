package usermdl

import (
	"gosplitwise/app/pkg/mongodbmdl"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestAddFriends(t *testing.T) {
	mongodbmdl.Init()
	ids := []primitive.ObjectID{}
	objectId, _ := primitive.ObjectIDFromHex("66463e4ccb2ea5bbec05def3")
	ids = append(ids, objectId)
	var want int64 = 1
	got, err := AddFriends("prajwal.u.patil@gmail.com", ids)
	if err != nil {
		t.Error(err)
	}
	if want == got {
		t.Log("success")
	} else {
		t.Errorf("want %d got %d", want, got)
	}
}
