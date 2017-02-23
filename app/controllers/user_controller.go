package controllers

import (
    "github.com/revel/revel"
    "myapp8/app/services"
)

type UserController struct {
    CommonController
}

func (c UserController) List() revel.Result {
    offset, limit, orderby, sort := c.validateQueryPram()
    return c.RenderJson(services.UserService{}.List(offset, limit, orderby, sort))
}	

func (c UserController) Get(user_id int64) revel.Result {
    return c.RenderJson(services.UserService{}.Get(user_id))
}	

func (c UserController) Create() revel.Result {
    return c.RenderJson(services.UserService{}.Create(c.Request.Body))
}	

func (c UserController) Update(user_id int64) revel.Result {
    return c.RenderJson(services.UserService{}.Update(c.Request.Body, user_id))
}	
