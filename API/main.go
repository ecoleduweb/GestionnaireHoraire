package main

import (
	"llio-api/routes"
	"github.com/gin-gonic/gin"
	"llio-api/auth"
)

func main() {
	auth.NewAuth()
	r := gin.Default()
	routes.RegisterRoutes(r)
	routes.AuthRoutes(r)
	r.Run(":8080")
}
