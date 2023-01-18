package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jjeanjacques10/crud-golang/src/configuration/logger"
	"github.com/jjeanjacques10/crud-golang/src/controller"
	"github.com/jjeanjacques10/crud-golang/src/controller/routes"
	"github.com/jjeanjacques10/crud-golang/src/model/service"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Init dependencies
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(
		service,
	)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
