package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KernelH132/pingme/internal/handler"
	"github.com/KernelH132/pingme/internal/repository"

	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}
}

func main() {
	repository.Connect()
	fmt.Println("Server running on :3000")
	http.ListenAndServe(":3000", http.HandlerFunc(handler.Handler))
}
