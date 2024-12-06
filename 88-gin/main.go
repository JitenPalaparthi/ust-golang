package main

import (
	"demo/database"
	"demo/handlers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	var dsn string

	if dsn = os.Getenv("DBCONN"); dsn == "" {
		dsn = "host=localhost user=admin password=admin123 dbname=demodb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	}

	db, err := database.GetConnection(dsn)
	if err != nil {
		log.Fatalln(err.Error())
	}

	userDB := database.NewUserDB(db)
	uhandler := handlers.NewUser(userDB)
	r := gin.Default()
	r.GET("/ping", handlers.Ping)
	r.GET("/health", handlers.Health)

	r.POST("/user", uhandler.CreateUser)
	r.GET("/user/:id", uhandler.GetUserByID)
	r.DELETE("/user/:id", uhandler.DeleteUserByID)
	r.PUT("/user/:id", uhandler.UpdateUserByID)
	if err := r.Run(); err != nil {
		log.Fatal(err.Error())
	} // default 8080
}
