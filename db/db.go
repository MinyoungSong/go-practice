package db

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"

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
	// clientOptions := options.Client().ApplyURI(dbURL)
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

}

func Select(filter interface{}) (*mongo.Cursor, error) {

	return collection.Find(context.TODO(), filter)

}

func SelectOne(filter interface{}) *mongo.SingleResult {

	return collection.FindOne(context.TODO(), filter)

}

func Insert() {

}

func Update() {

}

func Delete() {

}

// func Select(filter bson.D) []model.Clusterprovisions {

// 	// Connect to MongoDB
// 	clientOptions := options.Client().ApplyURI(dbURL)
// 	client, err := mongo.Connect(context.TODO(), clientOptions)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Check the connection
// 	err = client.Ping(context.TODO(), nil)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	collection = client.Database("iam").Collection("clusterprovisions")

// 	var results []model.Clusterprovisions
// 	var result model.Clusterprovisions

// 	collection.FindOne(context.TODO(), bson.D{{"metaData.clusterName", "native"}}).Decode(&result)

// 	results = append(results, result)

// 	return results

// }

// func SelectJSONData() kubeapi.Config {
// 	jsonData := []byte(`{"kind":"Config","apiVersion":"v1","clusters":[{"name":"native","cluster":{"server":"https://169.56.112.84:6443","insecure-skip-tls-verify":true}}],"users":[{"name":"admin","user":{"token":"eyJhbGciOiJSUzI1NiIsImtpZCI6IiJ9.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJ6Y3AtY2xpIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6InpjcC1jbGktYmFja2VuZC1zZXJ2aWNlLWFkbWluLXRva2VuLWdyd2s3Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6InpjcC1jbGktYmFja2VuZC1zZXJ2aWNlLWFkbWluIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQudWlkIjoiODg5ZTk0MGEtMDQ1NC0xMWVhLTkyY2MtMDY4YmRhZDQ4Nzc4Iiwic3ViIjoic3lzdGVtOnNlcnZpY2VhY2NvdW50OnpjcC1jbGk6emNwLWNsaS1iYWNrZW5kLXNlcnZpY2UtYWRtaW4ifQ.MCqNHYlXu4DolTsx4E-OTC4GWiJVX4IpAh6Zj9vdrmQfXFSrzL5JXsVYsPn5rivkFcC9Vcfp1yaar9bnNCGY0daZFHRa_04Ul7cJ3m1D-yeZSvEX17ClC47nfgEDFC-CFTysAeqTfHrf8yk-Ln4wBhbzsPahf2tL4eXDTzzvl-IWewIJPkkVpD5ad6UFcG45F5UfA-yVMaZxJTjszW9RJx3RegHQj5fIlQ_4_AuHioqvTkNC9yPGG5sa0F2nNZnkwAmbUI0OH6nMQQIoYo-ErYvlNgVDuI0hWJirDQ7TX1EXmeRjmFrk0m50QCY-px7EIBWEr-mYMen6J80yFkca8g"}}],"contexts":[{"name":"native-context","context":{"cluster":"native","user":"admin","namespace":"default"}}],"current-context":"native-context"}`)

// 	var kubeconfig kubeapi.Config
// 	var dat map[string]interface{}

// 	json.Unmarshal(jsonData, &kubeconfig)
// 	json.Unmarshal(jsonData, &dat)

// 	// fmt.Println(kubeconfig)
// 	// fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@")
// 	// fmt.Println(dat)

// 	// s, _ := json.Marshal(kubeconfig)

// 	// fmt.Println("$$$$$$$$$$")
// 	// fmt.Println(s)
// 	// fmt.Println("$$$$$$$$$$")
// 	// fmt.Println(string(s))

// 	return kubeconfig
// }
