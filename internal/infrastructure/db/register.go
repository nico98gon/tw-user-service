package db

import (
	"context"
	"user-service/internal/domain/users"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertRegister(u users.User) (string, bool, error) {
	ctx := context.TODO()

	db := MongoCN.Database(DatabaseName)
	col := db.Collection("users")

	u.Password, _ = EncryptPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil
}