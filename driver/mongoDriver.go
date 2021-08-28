package driver

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	Client *mongo.Client
}

var Mongo = &MongoDB{}

func ConnectMongoDB(user, password string) *MongoDB {

	connStr := getConnectionString(user, password)
	client, err := mongo.NewClient(options.Client().ApplyURI(connStr))

	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	fmt.Println("connection ok")
	Mongo.Client = client
	return Mongo
}

func (driver *MongoDB) ConnectCollection(databaseName, collectionName string) *mongo.Collection {
	return driver.Client.Database(databaseName).Collection(collectionName)
}

func getConnectionString(user, password string) string {

	if os.Getenv("LOCAL_MODE") == "on" {
		return os.Getenv("MONGODB_CONNECTION_LOCAL")
	}

	connStr := fmt.Sprintf(os.Getenv("MONGODB_CONNECTION_ONL"), user, password)

	return connStr
}
