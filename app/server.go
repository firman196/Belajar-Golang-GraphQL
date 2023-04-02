package app

import (
	"belajar-golang-gql/app/http"
	"belajar-golang-gql/database"
	"belajar-golang-gql/models"
	"belajar-golang-gql/utils"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Run() {
	//Load environtment
	envErr := godotenv.Load(".env")
	utils.SetFatalError(envErr)
	var dbUsername = os.Getenv("DB_USERNAME")
	var dbPassword = os.Getenv("DB_PASSWORD")
	var dbName = os.Getenv("DB_DATABASE")
	var dbHost = os.Getenv("DB_HOST")
	var dbPort = os.Getenv("DB_PORT")

	db, err := database.ConnectDB(dbUsername, dbPassword, dbHost, dbPort, dbName)
	if err != nil {
		utils.SetFatalError(err)
	}
	db.AutoMigrate(&models.User{})

	server := gin.Default()
	server.GET("/", http.PlaygroundHandler())
	server.POST("/query", http.GraphqlHandler())

}
