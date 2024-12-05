package main

import (
	"demo/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", handlers.Ping)
	r.GET("/health", handlers.Health)

	r.POST("/user", handlers.NewUser().Create)
	if err := r.Run(); err != nil {
		log.Fatal(err.Error())
	} // default 8080
}
