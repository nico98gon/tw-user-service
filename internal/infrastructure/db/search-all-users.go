package db

import (
	"context"
	"fmt"
	"user-service/internal/domain/relations"
	"user-service/internal/domain/users"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SearchAllUsers(ID string, cursor string, search string, typeUser string) ([]*users.User, string, bool) {
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("users")

	var results []*users.User

	filter := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}
	if cursor != "" {
		filter["_id"] = bson.M{"$lt": cursor}
	}

	options := options.Find()
	options.SetLimit(20)
	options.SetSort(bson.D{{Key: "name", Value: -1}})

	cur, err := col.Find(ctx, filter, options)
	if err != nil {
		fmt.Println("Error al buscar usuarios:", err.Error())
		return results, "", false
	}

	var include bool
	for cur.Next(ctx) {
		var u users.User

		err := cur.Decode(&u)
		if err != nil {
			fmt.Println("Error al decodificar usuario:", err.Error())
			return results, "", false
		}

		var rel relations.Relation
		rel.UserID = ID
		rel.UserIDRel = u.ID.Hex()

		include = false

		found := SearchRelation(rel)
		if typeUser == "new" && !found {
			include = true
		} else if typeUser == "follow" && found {
			include = true
		}

		if rel.UserIDRel == ID {
			include = false
		}
		if include {
			u.Password = ""
			u.Banner = ""
			u.Bio = ""
			u.Location = ""
			results = append(results, &u)
		}
	}

	if err := cur.Err(); err != nil {
		fmt.Println("Error al recorrer el cursor:", err.Error())
		return results, "", false
	}

	cur.Close(ctx)

	var nextCursor string
	if len(results) > 0 {
		nextCursor = results[len(results)-1].ID.Hex()
	}

	return results, nextCursor, true
}