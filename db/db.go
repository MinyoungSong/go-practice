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

	"github.com/Kamva/mgm"
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

	// Setup mgm default config
	err := mgm.SetDefaultConfig(nil, "iam", clientOptions)

	// // Connect to MongoDB
	// client, err := mongo.Connect(context.TODO(), clientOptions)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Check the connection
	// err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	// collection = client.Database("iam").Collection("clusterprovisions")

	fmt.Println("Connected to MongoDB!")

}

func Select(filter bson.D) []model.Clusterprovisions {

	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI(dbURL)
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

	var results []model.Clusterprovisions
	var result model.Clusterprovisions

	resultOne := collection.FindOne(context.TODO(), bson.D{{"metaData.clusterName", "native"}})
	resultOne.Decode(&result)
	rawByte, _ := resultOne.DecodeBytes()

	// err1 := bson.Unmarshal(rawByte, &result)
	// if err1 == nil {
	// 	fmt.Println(result)
	// }

	fmt.Println("########################")

	fmt.Println(string(rawByte))
	results = append(results, result)

	// cur, err := collection.Find(context.TODO(), filter)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Finding multiple documents returns a cursor
	// // Iterating through the cursor allows us to decode documents one at a time
	// for cur.Next(context.TODO()) {

	// 	// create a value into which the single document can be decoded
	// 	var elem model.Clusterprovisions

	// 	err := cur.Decode(&elem)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	results = append(results, elem)
	// }

	// if err := cur.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	// Close the cursor once finished
	// defer cur.Close(context.TODO())

	// fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)

	return results

}

func Insert() {

}

func Update() {

}

func Delete() {

}
