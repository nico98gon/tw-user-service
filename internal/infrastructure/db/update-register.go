package db

import (
	"context"
	"user-service/internal/domain/users"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateRegister(u users.User, ID string) (bool, error) {
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("users")

	register := make(map[string]interface{})
	if len(u.Name) > 0 {
		register["name"] = u.Name
	}
	if len(u.LastName) > 0 {
		register["lastName"] = u.LastName
	}
	if len(u.Password) > 0 {
		register["password"] = u.Password
	}
	if len(u.Bio) > 0 {
		register["bio"] = u.Bio
	}
	if len(u.WebSite) > 0 {
		register["webSite"] = u.WebSite
	}
	if len(u.Location) > 0 {
		register["location"] = u.Location
	}

	updateStr := bson.M{"$set": register}
	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filter, updateStr)
	if err != nil {
		return false, err
	}

	return true, nil
}