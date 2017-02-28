package controllers

import (
    "github.com/revel/revel"
    "github.com/syedomair/revel_api/app/models"
    "github.com/syedomair/revel_api/app/services"
    //"fmt"
)

type App struct {
    *revel.Controller
    services.CommonService
}

func (c App) Index() revel.Result {

    user := &models.User{Email: "Kiswono Prayogo", FirstName: "first name", LastName: "Last name"}
    c.Tx.NewRecord(user)
    c.Tx.Create(user)

    //user := []models.User{}
    //tuser := services.Db.Find(&user)
    //fmt.Println(tuser)

	return c.Render()
}
