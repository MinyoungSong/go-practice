package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"skcloud.io/cloudzcp/zcpctl-backend/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbURL string = ""
var collection *mongo.Collection

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func InitDB() {

	if os.Getenv("ZCP_CLI_BACKEND_VERSION") != "" {
		dbURL = os.Getenv("MONGODB_URL_K8s")
	} else {
		dbURL = os.Getenv("MONGODB_URL_LOCAL")
	}

	clientOptions := options.Client().ApplyURI(dbURL)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("iam").Collection("clusterprovisions")

	fmt.Println("Connected to MongoDB!")

}

func Select() {

	var results []model.ClusterSchema

	// err := collection.FindOne(context.TODO(), bson.D{{}}).Decode(&results)

	cur, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem model.ClusterSchema

		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	defer cur.Close(context.TODO())

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)

}

func Insert() {

}

func Update() {

}

func Delete() {

}
