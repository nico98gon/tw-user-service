package db

import (
	"context"
	"user-service/internal/domain/users"

	"go.mongodb.org/mongo-driver/bson"
)



func UserAlreadyExists(email string) (users.User, bool, string) {
	ctx := context.TODO()

	db := MongoCN.Database(DatabaseName)
	col := db.Collection("users")

	condition := bson.M{"email": email}
	var result users.User

	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
