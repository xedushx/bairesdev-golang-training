package api

import (
	"net/http"

	"bairesdev.com/golang/training/questionsandanswers/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcome)
	router.GET("/question", controllers.GetAllQuestions)
	router.GET("/question/user/:userId", controllers.GetAllQuestionsByUserId)
	router.POST("/question", controllers.CreateQuestion)
	router.GET("/question/:id", controllers.GetQuestionById)
	router.PUT("/question/:id", controllers.UpdateQuestion)
	router.DELETE("/question/:id", controllers.DeleteQuestion)
	router.NoRoute(notFound)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To Question and Answers API",
	})
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
	return
}
