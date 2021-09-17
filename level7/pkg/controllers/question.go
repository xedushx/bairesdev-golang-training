package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	guuid "github.com/google/uuid"
)

type Question struct {
	ID          string    `json:"id"`
	Summary     string    `json:"summary"`
	Description string    `json:"description"`
	Answer      string    `json:"answer"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	CreatedUser string    `json:"createdUser"`
	UpdatedAt   time.Time `json:"updatedAt"`
	UpdatedUser string    `json:"updatedUser"`
}

// Create Question Table
func CreateQuestionTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&Question{}, opts)
	if createError != nil {
		log.Printf("Error while creating Question table, Reason: %v\n", createError)
		return createError
	}
	log.Printf("Question table created")
	return nil
}

// INITIALIZE DB CONNECTION (TO AVOID TOO MANY CONNECTION)
var dbConnect *pg.DB

func InitiateDB(db *pg.DB) {
	dbConnect = db
}

func GetAllQuestions(c *gin.Context) {
	var questionList []Question

	err := dbConnect.Model(&questionList).Select()

	if err != nil {
		log.Printf("Error while getting all Questions, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Questions",
		"data":    questionList,
	})
	return
}

func GetAllQuestionsByUserId(c *gin.Context) {
	userId := c.Param("userId")
	var questionList []Question
	err := dbConnect.Model(&questionList).Where("created_user = ?", userId).Select()

	if err != nil {
		log.Printf("Error while getting all questions by user Id, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Questions not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Questions",
		"data":    questionList,
	})
	return
}

func GetQuestionById(c *gin.Context) {
	questionId := c.Param("id")
	question := &Question{ID: questionId}
	err := dbConnect.Select(question)

	if err != nil {
		log.Printf("Error while getting a single question, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Question not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single Question",
		"data":    question,
	})
	return
}

func CreateQuestion(c *gin.Context) {
	var question Question
	c.BindJSON(&question)
	id := guuid.New().String()
	summary := question.Summary
	description := question.Description
	answer := question.Answer
	status := "active"
	createdAt := time.Now()
	createdUser := question.CreatedUser
	updatedAt := time.Now()
	updatedUser := question.UpdatedUser

	insertError := dbConnect.Insert(&Question{
		ID:          id,
		Summary:     summary,
		Description: description,
		Answer:      answer,
		Status:      status,
		CreatedAt:   createdAt,
		CreatedUser: createdUser,
		UpdatedAt:   updatedAt,
		UpdatedUser: updatedUser,
	})
	if insertError != nil {
		log.Printf("Error while inserting new Question into db, Reason: %v\n", insertError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Question created Successfully",
	})
	return
}

func UpdateQuestion(c *gin.Context) {
	questionId := c.Param("id")
	var question Question
	c.BindJSON(&question)
	summary := question.Summary
	description := question.Description
	answer := question.Answer
	status := question.Status
	updatedAt := time.Now()
	updatedUser := question.UpdatedUser

	_, err := dbConnect.Model(&Question{}).Set(
		"summary = ?, description = ?, answer = ?, status = ?, updated_at = ?, updated_user = ?",
		summary, description, answer, status, updatedAt, updatedUser).Where("id = ?", questionId).Update()
	if err != nil {
		log.Printf("Error, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Question Updated Successfully",
	})
	return
}

func DeleteQuestion(c *gin.Context) {
	questionId := c.Param("id")
	question := &Question{ID: questionId}

	err := dbConnect.Delete(question)
	if err != nil {
		log.Printf("Error while deleting a single Question, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Question deleted successfully",
	})
	return
}
