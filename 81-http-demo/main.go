package main

import (
	"context"
	"demo/handlers"
	"demo/utils"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// args := os.Args
	// log.Println("Command line args:", args)
	//var PORT string
	PORT := os.Getenv("PORT")
	if PORT == "" {
		//PORT = "8083"
		flag.StringVar(&PORT, "port", "8083", "--port=8083 or -port=8083")
		flag.Parse()
	}
	// http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Write([]byte("pong"))
	// })

	parent := context.Background()

	ctx, cancel := signal.NotifyContext(parent, os.Interrupt, syscall.SIGKILL, syscall.SIGTERM)
	defer cancel()
	utils.FileName = "data/users.db"
	go utils.Write(ctx)

	http.HandleFunc("/ping", handlers.Ping)
	http.HandleFunc("/health", handlers.Health)
	http.HandleFunc("/user", handlers.CreateUser)

	log.Println("Service is up and running on port:", PORT)
	if err := http.ListenAndServe(":"+PORT, nil); err != nil { // 0.0.0.0:8083
		log.Fatalln(err.Error())
	}
}

// go does not require any kind container to deploy the web application. No addtional web server
