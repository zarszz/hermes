package utils

import "github.com/gin-gonic/gin"

type Response struct {
	StatusCode int         `json:"statusCode"`
	Method     string      `json:"method"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type ResponseError struct {
	StatusCode int         `json:"statusCode"`
	Method     string      `json:"method"`
	Error      interface{} `json:"error"`
}

func HandleResponse(c *gin.Context, message string, code int, method string, data interface{}) {
	c.JSON(code, Response{
		StatusCode: code,
		Method:     method,
		Message:    message,
		Data:       data,
	})
	if code >= 400 {
		defer c.AbortWithStatus(code)
	}
}

func HandleErrorResponse(c *gin.Context, code int, method string, err interface{}) {
	c.JSON(code, ResponseError{
		StatusCode: code,
		Method:     method,
		Error:      err,
	})
	defer c.AbortWithStatus(code)
}
