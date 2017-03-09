package controllers

import (
	"github.com/revel/revel"
)

type CommonController struct {
	*revel.Controller
}

func (c CommonController) validateQueryPram() (string, string, string, string) {
	offset := c.Params.Get("offset")
	limit := c.Params.Get("limit")
	orderby := c.Params.Get("orderby")
	sort := c.Params.Get("sort")
	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = "10000"
	}
	if orderby == "" {
		orderby = "id"
	}
	if sort == "" {
		sort = "asc"
	}
	return offset, limit, orderby, sort
}
