package controllers

import (
	"database/sql"
	//"fmt"
	"github.com/revel/revel"
)

//DB represents the database instance
var db *sql.DB

//InitDB initializes DB connection
func InitDB() {
	//connstring := fmt.Sprintf("user=%s password='%s' dbname=%s sslmode=disable", 
        //    revel.Config.StringDefault("db.user_name", ""), 
        //    revel.Config.StringDefault("db.password", ""), 
        //    revel.Config.StringDefault("db.db_name", ""))

	var err error
	//db, err = sql.Open("postgres", connstring)
        //db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
        db, err = sql.Open("postgres", revel.Config.StringDefault("db.spec", "") )
	if err != nil {
		revel.INFO.Println("DB Error", err)
	}
	revel.INFO.Println("DB Connected")
}

func init() {
	revel.OnAppStart(InitDB)
}
