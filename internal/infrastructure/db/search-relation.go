package db

import (
	"context"
	"user-service/internal/domain/relations"

	"go.mongodb.org/mongo-driver/bson"
)

func SearchRelation(rel relations.Relation) bool {
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("relation")

	condition := bson.M{"user_id": rel.UserID, "user_id_rel": rel.UserIDRel}

	var result relations.Relation

	err := col.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return false
	}

	return true
}