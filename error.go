package GinX

import "github.com/gin-gonic/gin"

func StringIfErrNotNil(c *gin.Context, err error, code int, response string) bool {
	if err == nil {
		return false
	}
	c.String(code, response)
	return false
}

func JsonIfErrNotNil(c *gin.Context, err error, code int, response interface{}) bool {
	if err == nil {
		return false
	}
	c.JSON(code, response)
	return false
}

func FileIfErrNotNil(c *gin.Context, err error, code int, path string) bool {
	if err == nil {
		return false
	}
	c.File(path)
	return false
}
