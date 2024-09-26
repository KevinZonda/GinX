package GinX

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Gin200Json(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func Gin200String(c *gin.Context, data string) {
	c.String(http.StatusOK, data)
}

func Gin200File(c *gin.Context, filepath string) {
	c.File(filepath)
}

func Gin301Redirect(c *gin.Context, url string) {
	c.Redirect(http.StatusMovedPermanently, url)
}

func Gin302Redirect(c *gin.Context, url string) {
	c.Redirect(http.StatusFound, url)
}

func Gin400Json(c *gin.Context, data interface{}) {
	c.JSON(http.StatusBadRequest, data)
}

func Gin400String(c *gin.Context, data string) {
	c.String(http.StatusBadRequest, data)
}

func Gin401Json(c *gin.Context, data interface{}) {
	c.JSON(http.StatusUnauthorized, data)
}

func Gin401String(c *gin.Context, data string) {
	c.String(http.StatusUnauthorized, data)
}

func Gin403Json(c *gin.Context, data interface{}) {
	c.JSON(http.StatusForbidden, data)
}

func Gin403String(c *gin.Context, data string) {
	c.String(http.StatusForbidden, data)
}

func Gin404Json(c *gin.Context, data interface{}) {
	c.JSON(http.StatusNotFound, data)
}

func Gin404String(c *gin.Context, data string) {
	c.String(http.StatusNotFound, data)
}

func Gin500Json(c *gin.Context, data interface{}) {
	c.JSON(http.StatusInternalServerError, data)
}

func Gin500String(c *gin.Context, data string) {
	c.String(http.StatusInternalServerError, data)
}

func Gin501Json(c *gin.Context, data interface{}) {
	c.JSON(http.StatusNotImplemented, data)
}

func Gin501String(c *gin.Context, data string) {
	c.String(http.StatusNotImplemented, data)
}

func Gin502Json(c *gin.Context, data interface{}) {
	c.JSON(http.StatusBadGateway, data)
}

func Gin502String(c *gin.Context, data string) {
	c.String(http.StatusBadGateway, data)
}

func Gin503Json(c *gin.Context, data interface{}) {
	c.JSON(http.StatusServiceUnavailable, data)
}

func Gin503String(c *gin.Context, data string) {
	c.String(http.StatusServiceUnavailable, data)
}

func Gin504Json(c *gin.Context, data interface{}) {
	c.JSON(http.StatusGatewayTimeout, data)
}

func Gin504String(c *gin.Context, data string) {
	c.String(http.StatusGatewayTimeout, data)
}

func Gin505Json(c *gin.Context, data interface{}) {
	c.JSON(http.StatusHTTPVersionNotSupported, data)
}

func Gin505String(c *gin.Context, data string) {
	c.String(http.StatusHTTPVersionNotSupported, data)
}

func Gin506Json(c *gin.Context, data interface{}) {
	c.JSON(http.StatusVariantAlsoNegotiates, data)
}

func Gin506String(c *gin.Context, data string) {
	c.String(http.StatusVariantAlsoNegotiates, data)
}

func GinString(c *gin.Context, code int, data string) {
	c.String(code, data)
}

func GinJson(c *gin.Context, code int, data interface{}) {
	c.JSON(code, data)
}
