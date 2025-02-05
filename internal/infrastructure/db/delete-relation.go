package db

import (
	"context"
	"user-service/internal/domain/relations"
)

func DeleteRelationFromDB(rel relations.Relation) (bool, error) {
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("relation")

	_, err := col.DeleteOne(ctx, rel)
	if err != nil {
		return false, err
	}

	return true, nil
}