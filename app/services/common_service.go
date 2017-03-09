package services

import (
	"github.com/revel/revel"
)

type CommonService struct {
	*revel.Controller

	response map[string]interface{}
	DatabaseService
}

func (c CommonService) errorAuthResponse(class interface{}) map[string]interface{} {
	return c.commonResponse(class, "error", "400")
}

func (c CommonService) errorResponse(class interface{}) map[string]interface{} {
	return c.commonResponse(class, "error", "500")
}

func (c CommonService) successResponse(class interface{}) map[string]interface{} {
	return c.commonResponse(class, "success", "200")
}

func (c CommonService) commonResponse(class interface{}, result string, code string) map[string]interface{} {
	c.response = make(map[string]interface{})
	c.response["data"] = class
	c.response["result"] = result
	c.response["code"] = code
	return c.response
}

func (c CommonService) successResponseList(class interface{}, offset string, limit string, count string) map[string]interface{} {
	tempResponse := make(map[string]interface{})
	tempResponse["offset"] = offset
	tempResponse["limit"] = limit
	tempResponse["count"] = count
	tempResponse["list"] = class
	return c.successResponse(tempResponse)
}
