package services

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/syedomair/revel_api/app/models"
	"io"
	"strconv"
	//"golang.org/x/crypto/scrypt"
	//"github.com/revel/revel"
)

type UserService struct {
	CommonService
}

type UserResponse struct {
	Id        int64  `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (c UserService) List(offset string, limit string, orderby string, sort string) map[string]interface{} {

	count := 0
	userResponse := []UserResponse{}
	Db.Table("public.user").
		Select(" * ").
		Count(&count).
		Limit(limit).
		Offset(offset).
		Order(orderby + " " + sort).
		Scan(&userResponse)

	return c.successResponseList(userResponse, offset, limit, strconv.Itoa(count))
}

func (c UserService) Get(userId int64) map[string]interface{} {

	userResponse := UserResponse{}
	Db.Table("public.user as u").
		Select("*").
		Where("u.id = ?", userId).
		Scan(&userResponse)

	return c.successResponse(userResponse)
}

func (c UserService) Create(jsonString io.Reader) map[string]interface{} {
	user := models.User{}

	decodedJson := json.NewDecoder(jsonString)
	var jsonMap map[string]interface{}

	if err := decodedJson.Decode(&jsonMap); err != nil {
		fmt.Println(err)
	}

	if val, ok := jsonMap["first_name"]; ok {
		user.FirstName = val.(string)
	} else {
		return c.errorResponse("first_name is a requird field")
	}

	if val, ok := jsonMap["last_name"]; ok {
		user.LastName = val.(string)
	} else {
		return c.errorResponse("last_name is a requird field")
	}

	if val, ok := jsonMap["email"]; ok {
		user.Email = val.(string)
	} else {
		return c.errorResponse("email is a requird field")
	}

	if val, ok := jsonMap["password"]; ok {
		user.Password = val.(string)
	} else {
		return c.errorResponse("password is a requird field")
	}

	plainPassword, _ := b64.StdEncoding.DecodeString(user.Password)
	//salt := revel.Config.StringDefault("server_salt", "")
	//encryptedPassword, _ := scrypt.Key([]byte(plainPassword), []byte(salt), 16384, 8, 1, 32)
	user.Password = string(plainPassword)

	Db.NewRecord(user)
	Db.FirstOrCreate(&user, user)
	fmt.Println(user.Id)

	return c.successResponse(user.Id)
}

func (c UserService) Update(jsonString io.Reader, userId int64) map[string]interface{} {
	user := models.User{}
	inputUser := models.User{}

	decodedJson := json.NewDecoder(jsonString)
	var jsonMap map[string]interface{}

	if err := decodedJson.Decode(&jsonMap); err != nil {
		fmt.Println(err)
	}

	if val, ok := jsonMap["first_name"]; ok {
		inputUser.FirstName = val.(string)
	}

	if val, ok := jsonMap["last_name"]; ok {
		inputUser.LastName = val.(string)
	}

	if val, ok := jsonMap["email"]; ok {
		inputUser.Email = val.(string)
	}

	Db.First(&user, userId)
	Db.Model(&user).Updates(&inputUser)

	return c.successResponse(userId)
}

func (c UserService) Authenticate(jsonString io.Reader) map[string]interface{} {

	decodedJson := json.NewDecoder(jsonString)
	var jsonMap map[string]interface{}

	if err := decodedJson.Decode(&jsonMap); err != nil {
		fmt.Println(err)
	}
	password, _ := b64.URLEncoding.DecodeString(jsonMap["password"].(string))

	//salt := revel.Config.StringDefault("server_salt", "")
	//encryptedPassword, _ := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, 32)

	userResponse := UserResponse{}
	Db.Table("public.user as u").
		Select("*").
		Where("email = ? AND password = ?", jsonMap["email"], password).
		Scan(&userResponse)

	return c.successResponse(userResponse)
}
