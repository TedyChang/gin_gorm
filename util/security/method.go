package security

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context, token string) {
	c.SetCookie("token", token, 60*60*24, "/", "localhost", true, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "login success",
	})
}
