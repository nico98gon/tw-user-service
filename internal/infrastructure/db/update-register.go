package db

import (
	"context"
	"fmt"
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
	if len(u.Avatar) > 0 {
		register["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		register["banner"] = u.Banner
	}

	updateStr := bson.M{"$set": register}
	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	updateResult, err := col.UpdateOne(ctx, filter, updateStr)
	if err != nil {
		return false, err
	}
	fmt.Printf("Número de documentos actualizados: %d\n", updateResult.ModifiedCount)

	if updateResult.ModifiedCount == 0 {
		return false, fmt.Errorf("Los datos a actualizar son iguales a los actuales")
	}

	return true, nil
}