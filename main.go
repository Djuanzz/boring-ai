package main

import (
	"fmt"

	"github.com/Djuanzz/boring-ai/config"
	"github.com/Djuanzz/boring-ai/routes"
)

func main() {
	fmt.Println("Boring AI")

	cfg := config.LoadEnv()

	server := routes.SetupRoutes(cfg)

	if err := server.Run(":" + cfg.Port); err != nil {
		fmt.Println("Error starting server:", err)
		panic(err.Error())
	}
}
