package main

import (
	"fmt"

	"github.com/Snehashish1609/validator-api/middlewares"

	"github.com/Snehashish1609/validator-api/models"

	v1 "github.com/Snehashish1609/validator-api/handlers/v1"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting Validator API...")

	// gin default router
	r := gin.Default()

	// latency middleware
	r.Use(middlewares.LatencyLogger())

	userHandler := models.NewUserHandler()
	apiHandler := v1.NewAPIHandler(userHandler)

	// routes
	r.POST("/validate-user", apiHandler.ValidateUser)

	// start server
	r.Run(":8080")
}
