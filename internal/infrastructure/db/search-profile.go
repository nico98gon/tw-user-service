package db

import (
	"context"
	"user-service/internal/domain/users"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SearchProfile(ID string) (users.User, error) {
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("users")

	var profile users.User
	var err error

	ObjID, objIDErr := primitive.ObjectIDFromHex(ID)
	if objIDErr == nil {
		condition := bson.M{"_id": ObjID}
		err = col.FindOne(ctx, condition).Decode(&profile)
	}

	if objIDErr != nil || err != nil {
		condition := bson.M{"_id": ID}
		err = col.FindOne(ctx, condition).Decode(&profile)
	}

	if err != nil {
		return profile, err
	}

	profile.Password = ""
	return profile, nil
}
