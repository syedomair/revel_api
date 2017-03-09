package services

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/revel/revel"
	"github.com/syedomair/revel_api/app/models"
)

type DatabaseService struct {
	Tx *gorm.DB
}

var Db *gorm.DB

func InitDB() {
	var err error
	Db, err = gorm.Open("postgres", revel.Config.StringDefault("db.spec", ""))

	if err != nil {
		revel.ERROR.Println("FATAL", err)
		panic(err)
	}
	Db.SingularTable(true)

	Db.DB().SetMaxIdleConns(revel.Config.IntDefault("db.max.idle.conns", 3))
	Db.DB().SetMaxOpenConns(revel.Config.IntDefault("db.max.open.conns", 10))
	Db.LogMode(revel.Config.BoolDefault("db.log.mode", true))
	Db.SetLogger(gorm.Logger{revel.INFO})

	//The following lines are needed when Heroku drops all the tables periodically.
	if !Db.HasTable(&models.Book{}) {
		Db.CreateTable(&models.Book{})
	}
	if !Db.HasTable(&models.Client{}) {
		Db.CreateTable(&models.Client{})
		client := models.Client{Name: "Test",
			ApiKey:    "dHb%e@Bg0f8-API_KEY-&bE71jKoH=2",
			ApiSecret: "g$5%6kQ56-API_SECRET-6gE@7&EbR2",
			Active:    true}
		Db.NewRecord(client)
		Db.Create(&client)
	}
	if !Db.HasTable(&models.User{}) {
		Db.CreateTable(&models.User{})
	}
}

func (c *DatabaseService) Begin() revel.Result {
	txn := Db.Begin()
	if txn.Error != nil {
		panic(txn.Error)
	}
	c.Tx = txn
	revel.INFO.Println("c.Tx init", c.Tx)
	return nil
}

func (c *DatabaseService) Commit() revel.Result {
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

func (c *DatabaseService) Rollback() revel.Result {
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
