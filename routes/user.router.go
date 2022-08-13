package routes

import (
	"gin_gorm/domain/user"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {
	r.GET("/user", user.GetName)
	r.POST("/user", user.PostSave)
	r.POST("/login", user.PostLogin)
}
