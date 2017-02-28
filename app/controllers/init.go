package controllers

import (
	"github.com/revel/revel"
)

func init() {
    revel.InterceptMethod(BookController.Init, revel.BEFORE)
    revel.InterceptMethod(UserController.Init, revel.BEFORE)
}
