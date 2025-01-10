package main

import (
	"fmt"

	"github.com/Snehashish1609/validator-api/config"
	v1 "github.com/Snehashish1609/validator-api/handlers/v1"
	"github.com/Snehashish1609/validator-api/middlewares"
	"github.com/Snehashish1609/validator-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting Validator API...")

	c := config.InitConfig("Validator", ":8080")

	// gin default router
	r := gin.Default()

	// latency middleware
	r.Use(middlewares.LatencyLogger())

	userHandler := models.NewUserHandler()
	apiHandler := v1.NewAPIHandler(c, userHandler)

	// routes
	r.POST("/validate-user", apiHandler.ValidateUser)

	// start server
	r.Run(c.Port)
}
