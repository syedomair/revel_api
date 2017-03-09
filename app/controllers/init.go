package controllers

import (
	"github.com/revel/revel"
	"github.com/syedomair/revel_api/app/services"
)

func init() {
	revel.InterceptMethod(BookController.Init, revel.BEFORE)
	revel.InterceptMethod(UserController.Init, revel.BEFORE)

	revel.OnAppStart(services.InitDB)
	revel.InterceptMethod((*services.DatabaseService).Begin, revel.BEFORE)
	revel.InterceptMethod((*services.DatabaseService).Commit, revel.AFTER)
	revel.InterceptMethod((*services.DatabaseService).Rollback, revel.FINALLY)
}
