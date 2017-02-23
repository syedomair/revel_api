package controllers

import (
    "github.com/revel/revel"
    "myapp8/app/services"
)

type BookController struct {
    CommonController
}

func (c BookController) List() revel.Result {
    offset, limit, orderby, sort := c.validateQueryPram()
    return c.RenderJson(services.BookService{}.List(offset, limit, orderby, sort))
}	

func (c BookController) Get(book_id int64) revel.Result {
    return c.RenderJson(services.BookService{}.Get(book_id))
}	

func (c BookController) Create() revel.Result {
    return c.RenderJson(services.BookService{}.Create(c.Request.Body))
}	

func (c BookController) Update(book_id int64) revel.Result {
    return c.RenderJson(services.BookService{}.Update(c.Request.Body, book_id))
}	
