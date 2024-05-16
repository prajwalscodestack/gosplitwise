package usermdl

import (
	"context"
	"gosplitwise/app/models"
	"gosplitwise/app/pkg/mongodbmdl"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Register(user models.User) error {
	DAO, err := mongodbmdl.NewMongoDB("gosplitwise", context.TODO())
	if err != nil {
		return err
	}
	_, saveErr := DAO.SaveDocument("user", user)
	if saveErr != nil {
		return saveErr
	}
	return saveErr
}

func Login(cred models.Credential) (*models.User, error) {
	DAO, err := mongodbmdl.NewMongoDB("gosplitwise", context.TODO())
	if err != nil {
		return nil, err
	}
	result := DAO.FetchDocument("user", bson.M{
		"email":    cred.Email,
		"password": cred.Password,
	})
	if result.Err() == mongo.ErrNoDocuments {
		return nil, mongo.ErrNoDocuments
	}
	var user models.User
	if err := result.Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func AddFriends(email string, id []primitive.ObjectID) (int64, error) {
	DAO, err := mongodbmdl.NewMongoDB("gosplitwise", context.TODO())
	if err != nil {
		return 0, err
	}
	res, err := DAO.UpdateDocument("user", bson.M{"email": email}, bson.M{
		"$push": bson.M{
			"friends": bson.M{"$each": id},
		},
	})
	if err != nil {
		return 0, err
	}
	return res.ModifiedCount, nil
}
