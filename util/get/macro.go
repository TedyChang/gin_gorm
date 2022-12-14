package get

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ok(c *gin.Context, s string) {
	c.JSON(http.StatusOK, gin.H{
		"message": s,
	})
}
func OkId(c *gin.Context, s uint) {
	c.JSON(http.StatusOK, gin.H{
		"id": s,
	})
}
func Bad(c *gin.Context, s string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"message": s,
	})
}
