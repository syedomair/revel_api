package services

import (
  "database/sql"
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
  "github.com/revel/revel"
)

type CommonService struct{
    *revel.Controller
    Tx *gorm.DB
    response map[string]interface{} 
}

var Db *gorm.DB

func (c CommonService) errorResponse(class interface{})map[string]interface{} {
    c.response = make(map[string]interface{}) 
    c.response["error_message"] = class 
    c.response["result"] = "error"
    c.response["code"] = "500"
    return c.response
}	

func (c CommonService) successResponse(class interface{})map[string]interface{} {
    c.response = make(map[string]interface{}) 
    c.response["data"] = class 
    c.response["result"] = "success"
    c.response["code"] = "200"
    return c.response
}	

func (c CommonService) successResponseList(class interface{}, offset string, limit string, count string)map[string]interface{} {
    tempResponse := make(map[string]interface{}) 
    tempResponse["offset"] = offset 
    tempResponse["limit"] = limit
    tempResponse["count"] = count
    tempResponse["list"] = class
    return c.successResponse(tempResponse)
}	


func InitDB() {
  var err error
  Db, err = gorm.Open("postgres", revel.Config.StringDefault("db.spec", "") )

  if err != nil {
    revel.ERROR.Println("FATAL", err)
    panic(err)
  }
  Db.SingularTable(true)
}


func (c *CommonService) Begin() revel.Result {
  txn := Db.Begin()
  if txn.Error != nil {
    panic(txn.Error)
  }
  c.Tx = txn
  revel.INFO.Println("c.Tx init", c.Tx)
  return nil
}

func (c *CommonService) Commit() revel.Result {
  if c.Tx == nil {
    return nil
  }
  c.Tx.Commit()
  if err := c.Tx.Error; err != nil && err != sql.ErrTxDone {
    panic(err)
  }
  c.Tx = nil
  revel.INFO.Println("c.Tx commited (nil)")
  return nil
}

func (c *CommonService) Rollback() revel.Result {
  if c.Tx == nil {
    return nil
  }
  c.Tx.Rollback()
  if err := c.Tx.Error; err != nil && err != sql.ErrTxDone {
    panic(err)
  }
  c.Tx = nil
  return nil
}

