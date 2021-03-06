package controllers

import (
	"github.com/revel/revel"
	"github.com/syedomair/revel_api/app/services"
)

type BookController struct {
	CommonController
}

func (c BookController) Init() revel.Result {
	authService := services.SecurityService{}.AuthProvider(c.Request)
	if authService == nil {
		return nil
	} else {
		return c.RenderJson(authService)
	}
}

func (c BookController) List() revel.Result {
	offset, limit, orderby, sort := c.validateQueryPram()
	return c.RenderJson(services.BookService{}.List(offset, limit, orderby, sort))
}

func (c BookController) PublicBooks() revel.Result {
	offset, limit, orderby, sort := c.validateQueryPram()
	return c.RenderJson(services.BookService{}.List(offset, limit, orderby, sort))
}

func (c BookController) MyBooks(user_id int64) revel.Result {
	offset, limit, orderby, sort := c.validateQueryPram()
	return c.RenderJson(services.BookService{}.MyBooks(user_id, offset, limit, orderby, sort))
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
