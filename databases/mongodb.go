package databases

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client is the MongoDB client
var Client *mongo.Client

// ConnectDB will connect the application with database via server URI
func ConnectDB(uri string) error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).
		SetServerAPIOptions(serverAPI)

	// Create MongoDB client then connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return err
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").
		RunCommand(
			context.TODO(),
			bson.D{{"ping", 1}}).
		Err(); err != nil {
		return err
	}
	fmt.Println("Pinged the deployment. Successfully connected to MongoDB!")
	Client = client
	return nil
}
