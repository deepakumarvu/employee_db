package main

import (
	"employee/service/router"
)

func main() {
	// Create router
	router := router.NewRouter()

	// Start server
	router.Run("localhost:8080")
}
