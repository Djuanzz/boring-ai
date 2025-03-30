package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	fmt.Println("Boring AI")

	server := gin.Default()

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "5000"
	}

	if err := server.Run(":" + port); err != nil {
		fmt.Println("Error starting server:", err)
		panic(err.Error())
	}

	fmt.Println("Server started on port", port)
}
