package db

import (
	"context"
	"user-service/internal/domain/relations"
	"user-service/internal/domain/users"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SearchFollowers(userID string) ([]users.User, error) {
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("relation")

	filter := bson.M{"user_id_rel": userID}
	cursor, err := col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var relations []relations.Relation
	if err = cursor.All(ctx, &relations); err != nil {
		return nil, err
	}

	var userIDs []primitive.ObjectID
	for _, rel := range relations {
		if id, err := primitive.ObjectIDFromHex(rel.UserID); err == nil {
			userIDs = append(userIDs, id)
		}
	}

	if len(userIDs) == 0 {
		return []users.User{}, nil
	}

	usersCol := db.Collection("users")
	filterUsers := bson.M{"_id": bson.M{"$in": userIDs}}
	projection := bson.M{"password": 0}
	opts := options.Find().SetProjection(projection)
	cursorUsers, err := usersCol.Find(ctx, filterUsers, opts)
	if err != nil {
		return nil, err
	}
	defer cursorUsers.Close(ctx)

	var users []users.User
	if err = cursorUsers.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}