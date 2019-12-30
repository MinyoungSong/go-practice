package auth

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

var keycloakServer string
var createTokenURL string
var verifyTokenURL string

var keycloakClientID string
var keycloakClientPassword string

var client *resty.Client

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	keycloakServer = os.Getenv("AUTH_SERVER_URL")
	createTokenURL = os.Getenv("AUTH_CREATE_TOKEN_URL")
	verifyTokenURL = os.Getenv("AUTH_VERIFY_TOKEN_URL")
	keycloakClientID = os.Getenv("KEYCLOAK_CLIENT_ID")
	keycloakClientPassword = os.Getenv("KEYCLOAK_CLIENT_PASSWORD")

	client = resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	client.SetHostURL(keycloakServer)
	client.SetHeaders(map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	})
	client.SetBasicAuth(keycloakClientID, keycloakClientPassword)

}

func CreateToken(user User) string {

	var token string = ""

	createTokenBody := map[string]string{
		"grant_type": "password",
		"username":   user.Username,
		"password":   user.Password,
	}

	client.SetFormData(createTokenBody)

	resp, err := client.R().Post(createTokenURL)

	if err != nil {
		log.Printf("resp Error!!!!")
	}

	var res = make(map[string]interface{})
	json.Unmarshal(resp.Body(), &res)
	token = fmt.Sprintf("%v", res["access_token"])

	return token

}
