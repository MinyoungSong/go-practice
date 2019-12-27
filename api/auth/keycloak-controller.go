package auth

import (
	"crypto/tls"
	"encoding/json"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
)

var keycloakServer = os.Getenv("AUTH_SERVER_URL")
var createTokenURL = os.Getenv("AUTH_CREATE_TOKEN_URL")
var verifyTokenURL = os.Getenv("AUTH_VERIFY_TOKEN_URL")

var keycloakClientID = os.Getenv("KEYCLOAK_CLIENT_ID")
var keycloakClientPassword = os.Getenv("KEYCLOAK_CLIENT_PASSWORD")

func CreateToken(user User) string {

	var token string = "aaaa"

	createTokenBody := map[string]string{
		"grant_type": "password",
		"username":   user.Username,
		"password":   user.Password,
	}

	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	client.SetHostURL(keycloakServer)
	client.SetHeaders(map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	})
	client.SetBasicAuth(keycloakClientID, keycloakClientPassword)
	client.SetFormData(createTokenBody)

	resp, err := client.R().Post(createTokenURL)

	if err != nil {
		echo.Logger.Debug(err)
	}

	result := json.Unmarshal(resp.Body())

	echo.Logger.Debug(result)

	return token

}
