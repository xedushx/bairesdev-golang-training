package config

import (
	"log"
	"os"

	"bairesdev.com/golang/training/questionsandanswers/pkg/controllers"
	"github.com/go-pg/pg"
)

// Connecting to db
func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "postgres",
		Password: "postgres",
		Addr:     "localhost:5432",
		Database: "training_q",
	}
	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")

	controllers.CreateQuestionTable(db)
	controllers.InitiateDB(db)

	return db
}
