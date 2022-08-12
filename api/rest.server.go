package api

import (
	"gin_gorm/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func RestApi() {
	r := gin.Default()

	routes.Router(r)

	err := r.Run()
	if err != nil {
		log.Printf("@@@RestApi r.Run error@@@\nerr : \n%v", err)
	}
}
