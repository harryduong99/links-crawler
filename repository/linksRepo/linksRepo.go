package linksRepo

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"song-chord-crawler/config"
	"song-chord-crawler/driver"
	"song-chord-crawler/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func StoreLink(link models.Link) error {
	collection := driver.Mongo.ConnectCollection(config.DB_NAME, config.COL_LINKS)

	bbytes, _ := bson.Marshal(link)
	_, err := collection.InsertOne(context.Background(), bbytes)

	if err != nil {
		return err
	}

	return nil
}

func IsLinkExist(href string) bool {
	link := models.Link{}

	collection := driver.Mongo.ConnectCollection(config.DB_NAME, config.COL_LINKS)
	result := collection.FindOne(context.Background(), bson.M{"url": href})

	err := result.Decode(&link)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return false
		}
	}

	return true
}

func All() []models.Link {
	var links []models.Link
	collection := driver.Mongo.ConnectCollection(config.DB_NAME, config.COL_LINKS)

	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		var link models.Link
		// & character returns the memory address of the following variable.
		err := cur.Decode(&link) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		links = append(links, link)
	}

	return links
}

func PrintAll() {
	collection := driver.Mongo.ConnectCollection(config.DB_NAME, config.COL_LINKS)

	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		fmt.Println("touch collection went wrong!")
	}
	fmt.Print(cur)

	for cur.Next(context.TODO()) {

		// declare a result BSON object
		var result bson.M
		err := cur.Decode(&result)

		// If there is a cursor.Decode error
		if err != nil {
			fmt.Println("cursor.Next() error:", err)
			os.Exit(1)

			// If there are no cursor.Decode errors
		} else {
			fmt.Println("\nresult type:", reflect.TypeOf(result))
			fmt.Println("result:", result)
		}
	}
}
