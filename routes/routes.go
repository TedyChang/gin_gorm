package routes

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	UserRouter(r)
}
