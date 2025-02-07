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

func ConnectMongo(ctx context.Context) error {
	user := ctx.Value(domain.Key("user")).(string)
	password := ctx.Value(domain.Key("password")).(string)
	host := ctx.Value(domain.Key("host")).(string)
	isSrv := ctx.Value(domain.Key("isSrv")).(bool)

	var connStr string
	if isSrv {
		connStr = fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", user, password, host)
	} else {
		connStr = fmt.Sprintf("mongodb://%s:%s@%s/?retryWrites=true&w=majority&authSource=admin", user, password, host)
	}

	fmt.Println("URI de conexión:", connStr)

	clientOptions := options.Client().ApplyURI(connStr)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return fmt.Errorf("error al conectar a MongoDB: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return fmt.Errorf("error al hacer ping a MongoDB: %v", err)
	}

	fmt.Println("Conexión exitosa a MongoDB")
	MongoCN = client
	DatabaseName = ctx.Value(domain.Key("database")).(string)
	return nil
}

func DBConnected() bool {
	if MongoCN == nil {
		return false
	}
	err := MongoCN.Ping(context.TODO(), nil)
	return err == nil
}