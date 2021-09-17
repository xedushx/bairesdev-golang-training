package main

import (
	"log"

	"bairesdev.com/golang/training/questionsandanswers/pkg/api"
	config "bairesdev.com/golang/training/questionsandanswers/pkg/configs"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect DB
	config.Connect()

	// Init Router
	router := gin.Default()

	// Route Handlers / Endpoints
	api.Routes(router)

	log.Fatal(router.Run(":4747"))

}
