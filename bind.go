package GinX

import "github.com/gin-gonic/gin"

func BindJSONFailedJson(c *gin.Context, target any, code int, json any) bool {
	err := c.ShouldBindJSON(target)
	if err == nil {
		return false
	}
	c.JSON(code, json)
	return true
}

func BindJSONFailedString(c *gin.Context, target any, code int, str string) bool {
	err := c.ShouldBindJSON(target)
	if err == nil {
		return false
	}
	c.String(code, str)
	return true
}

func BindJSONFailedFx(c *gin.Context, target any, fx func(*gin.Context, error)) bool {
	err := c.ShouldBindJSON(target)
	if err == nil {
		return false
	}
	if fx != nil {
		fx(c, err)
	}
	return true
}

func BindJSON_T[T any](c *gin.Context, code int, json any) (t T, err error) {
	err = c.ShouldBindJSON(&t)
	return
}
