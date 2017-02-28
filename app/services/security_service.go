package services

import (
    "myapp8/app/models"
    "github.com/dgrijalva/jwt-go"
    b64 "encoding/base64"
    "github.com/revel/revel"
    "strings"
    //"golang.org/x/crypto/scrypt"
)
type SecurityService struct{
    CommonService
}

func (c SecurityService) AuthProvider(request *revel.Request) (map[string]interface{}) {

    apiKey := request.Header.Get("x-key")
    jwtToken := request.Header.Get("x-jwt")

    publicEndPoint := false
    if(strings.Contains(request.URL.Path, "/public/")){
        publicEndPoint = true
    }

    if(apiKey == ""){
        return c.errorAuthResponse("Header missing: x-key ")
    }
    if(jwtToken == ""){
        return c.errorAuthResponse("Header missing: x-jwt ")
    }

    client := models.Client{}
    Db.Table("public.client as c").
        Select("*").
        Where("c.api_key = ?", apiKey).
        Scan(&client)

    if(client.ApiSecret == ""){
        return c.errorAuthResponse("Invalid api_key ")
    }

    type Claims struct {
        Username string `json:"username"`
        Password string `json:"password"`
        jwt.StandardClaims
    }

    tokenClaims := Claims{}

    _, err := jwt.ParseWithClaims(jwtToken, &tokenClaims, func(token *jwt.Token) (interface{}, error) {
        return []byte(client.ApiSecret), nil
    })

    if(err != nil){
        return c.errorAuthResponse("Invalid JWT Signature")
    }

    if(!publicEndPoint){
        if((tokenClaims.Username != "new_registration") && (tokenClaims.Password != "new_registration")){
            user := models.User{}
            password, _ := b64.URLEncoding.DecodeString(tokenClaims.Password)
            plainPassword := string(password)

            //salt := revel.Config.StringDefault("server_salt", "")  
            //encryptedPassword, _ := scrypt.Key([]byte(plainPassword), []byte(salt), 16384, 8, 1, 32)
            Db.Where("email = ? AND password = ?", tokenClaims.Username, string(plainPassword)).Find(&user)

            if(user.FirstName == ""){
                return c.errorAuthResponse("Invalid Email or Password ")
            }
        }
    }

    return nil
}

