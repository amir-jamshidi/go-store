package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetCookie(c *gin.Context, time int, key, value string) {
	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie(
		key,
		value,
		time*60*60,
		"/",
		"",
		false,
		true,
	)
}

func ClearCookie(c *gin.Context, key string) {
	c.SetCookie(
		key,
		"",
		-1,
		"/",
		"",
		false,
		true,
	)
}

func GetFromCookie(c *gin.Context, key string) (string, error) {
	value, err := c.Cookie(key)
	return value, err
}
