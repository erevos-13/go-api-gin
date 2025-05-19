package main

import (
	"example.com/gin-api/db"
	"example.com/gin-api/middlewares"
	"example.com/gin-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Code
	db.InitDB()
	server := gin.Default()

	groupEvent := server.Group("/", middlewares.AuthMiddleware)

	groupEvent.GET("/event", routes.GetEvent)
	groupEvent.GET("/event/:id", routes.GetEventById)
	groupEvent.POST("/event", routes.CreateEvent)

	groupEvent.PUT("/event/:id", routes.UpdateEvent)

	groupEvent.DELETE("/event/:id", routes.DeleteEvent)

	groupEvent.POST("/event/:id/register", routes.RegisterForEvent)
	groupEvent.DELETE("/event/:id/register", routes.CancelRegistration)

	server.POST("/user/signup", routes.SignUp)
	server.POST("/user/login", routes.SignIn)

	server.Run(":3000")

}
