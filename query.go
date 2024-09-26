package GinX

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func QueryTryGetInt(c *gin.Context, query string) (val int, ok bool) {
	var str string
	str, ok = c.GetQuery(query)
	if !ok {
		return
	}
	val, err := strconv.Atoi(str)
	ok = err == nil
	return
}

func QueryTryGetIntDefault(c *gin.Context, query string, defaultVal int) int {
	val, ok := QueryTryGetInt(c, query)
	if !ok {
		return defaultVal
	}
	return val
}
