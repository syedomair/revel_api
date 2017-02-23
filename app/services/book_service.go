package services

import (
    "myapp8/app/models"
    "strconv"
    "encoding/json"
    "io"
    "fmt"
)
type BookService struct{
    CommonService
}

type BookResponse struct {
    Id              int64  `json:"id"`
    UserId          int64  `json:"user_id"`
    FirstName       string `json:"first_name"`
    LastName        string `json:"last_name"`
    Name            string  `json:"book_name" `
    Description     string  `json:"description" `
    Publish         bool  `json:"publish" `
} 

func (c BookService) List(offset string, limit string, orderby string, sort string) map[string]interface{} {

    count := 0 
    bookResponse :=[]BookResponse{}
    Db.Table("book").
        Select("*").
        Joins("join public.user as u on book.user_id = u.id").
        Count(&count).
        Limit(limit).
        Offset(offset).
        Order(orderby +" "+ sort).
        Scan(&bookResponse)

    return c.successResponseList(bookResponse, offset, limit, strconv.Itoa(count))
}  

func (c BookService) Get(bookId int64) map[string]interface{} {

    bookResponse :=BookResponse{}
    Db.Table("book").
        Select("*").
        Joins("join public.user as u on book.user_id = u.id").
        Where("book.id = ?", bookId).
        Scan(&bookResponse)

    return c.successResponse(bookResponse)
}  

func (c BookService) Create(jsonString io.Reader) map[string]interface{} {
    book := models.Book{}

    decodedJson := json.NewDecoder(jsonString)
    var jsonMap map[string]interface{}

    if err := decodedJson.Decode(&jsonMap); err != nil {
        fmt.Println(err)
    }
    if val, ok := jsonMap["name"]; ok {
        book.Name = val.(string)
    }else{
        return c.errorResponse("name is a requird field")
    }

    if val, ok := jsonMap["description"]; ok {
        book.Description = val.(string)
    }else{
         return c.errorResponse("description is a requird field")
    }

    if val, ok := jsonMap["user_id"]; ok {
        book.UserId = int64(val.(float64))
    }else{
        return c.errorResponse("user_id is a requird field")
    }

    if val, ok := jsonMap["publish"]; ok {
        book.Publish = val.(bool)
    }

    Db.NewRecord(book)
    Db.Create(&book)

    return c.successResponse(book.Id)
}  

func (c BookService) Update(jsonString io.Reader, bookId int64) map[string]interface{} {
    book := models.Book{}
    inputBook := models.Book{}
    
    decodedJson := json.NewDecoder(jsonString)
    var jsonMap map[string]interface{}

    if err := decodedJson.Decode(&jsonMap); err != nil {
        fmt.Println(err)
    }
    if val, ok := jsonMap["name"]; ok {
        inputBook.Name = val.(string)
    }

    if val, ok := jsonMap["description"]; ok {
        inputBook.Description = val.(string)
    }

    if val, ok := jsonMap["publish"]; ok {
        inputBook.Publish = val.(bool)
    }

    Db.First(&book, bookId)
    Db.Model(&book).Updates(&inputBook)

    return c.successResponse(book.Id)
}  
