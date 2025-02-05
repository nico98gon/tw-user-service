package db

import (
	"context"
	"errors"
	"user-service/internal/domain/relations"
)

func InsertRelation(rel relations.Relation) (bool, error) { 
	if SearchRelation(rel) {
		return false, errors.New("la relación ya existe")
	}

	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("relation")

	_, err := col.InsertOne(ctx, rel)
	if err != nil {
		return false, err
	}

	return true, nil
}