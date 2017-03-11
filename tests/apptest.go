package tests

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/revel/revel/testing"
)

type AppTest struct {
	testing.TestSuite
	Apikey    string
	ApiSecret string
	UserName  string
	Password  string
}

func (t *AppTest) Before() {
	println("Set up")
	t.Apikey = "dHb%e@Bg0f8-API_KEY-&bE71jKoH=2"
	t.ApiSecret = "g$5%6kQ56-API_SECRET-6gE@7&EbR2"
	t.UserName = "omair5@gmail.com"
	t.Password = "1"

}

func (t *AppTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *AppTest) TestThatBooksWorks() {

	fmt.Println(t.CallGetBackend("GET", "/books"))
	t.AssertOk()
}

func (t *AppTest) After() {
	println("Tear down")
}

func (t *AppTest) CallGetBackend(methodType string, url string) map[string]interface{} {
	return t.CallBackend(methodType, url, []byte(""), false)
}

func (t *AppTest) CallBackend(methodType string, url string, jsonStr []byte, newuser bool) map[string]interface{} {

	var userName string
	var password string

	if newuser {
		userName = "new_registration"
		password = "new_registration"
	} else {
		userName = t.UserName
		password = t.Password
	}

	signedJwtToken := t.CreateJWTToken(userName, password)

	completeURL := t.BaseUrl() + url
	request := &testing.TestRequest{}
	switch methodType {
	case "GET":
		request = t.GetCustom(completeURL)
	case "POST":
		request = t.PostCustom(completeURL, "application/json", bytes.NewBuffer(jsonStr))
	}
	request.Header.Set("x-key", t.Apikey)
	request.Header.Set("x-jwt", signedJwtToken)
	request.Header.Set("Content-Type", "application/json")
	request.Send()

	bodyInterface := make(map[string]interface{})
	json.Unmarshal(t.ResponseBody, &bodyInterface)

	return bodyInterface
}

func (t *AppTest) CreateJWTToken(userName string, password string) string {
	apiSecret := t.ApiSecret
	signingKey := []byte(apiSecret)

	type Claims struct {
		Username string `json:"username"`
		Password string `json:"password"`
		jwt.StandardClaims
	}

	claims := Claims{
		userName,
		b64.StdEncoding.EncodeToString([]byte(password)),
		jwt.StandardClaims{
			Issuer: "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedJwtToken, _ := token.SignedString(signingKey)

	return signedJwtToken
}
