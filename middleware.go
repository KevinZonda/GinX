package GinX

import "github.com/gin-gonic/gin"

func MidTracer[T any](c *gin.Context, key string, nonceGenerator func() T) gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Set(key, nonceGenerator())
		c.Next()
	}
}

func MidStringTracer(c *gin.Context, key string, nonceGenerator func() string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(key, nonceGenerator())
		c.Next()
	}
}

func GetTraceId[T any](c *gin.Context, key string) T {
	var t T
	v, ok := c.Get(key)
	if !ok {
		return t
	}
	return v.(T)
}

func GetStringTraceId(c *gin.Context, key string) string {
	v, ok := c.Get(key)
	if !ok {
		return ""
	}
	return v.(string)
}
