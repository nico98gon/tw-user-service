package db

import (
	"context"
	"fmt"
	"user-service/internal/domain"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN *mongo.Client
var DatabaseName string

func ConnectMongo(ctx context.Context) (error) {
	user := ctx.Value(domain.Key("user")).(string)
	password := ctx.Value(domain.Key("password")).(string)
	host := ctx.Value(domain.Key("host")).(string)
	connStr := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", user, password, host)

	var clientOptions = options.Client().ApplyURI(connStr)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("Error en la conexión a la base de datos:", err)
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("Error en la conexión a la base de datos:", err)
		return err
	}

	fmt.Println("Conexión a la base de datos MongoDB exitosa")

	MongoCN = client
	DatabaseName = ctx.Value(domain.Key("database")).(string)

	return nil
}

func DBConnected() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	return err == nil
}