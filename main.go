package main

import (
	"fmt"
	"notes-api/routes"
)

func main() {
	// Welcome message printed when the server starts
	fmt.Println("🚀 Welcome! Notes API server is starting...")

	r := routes.SetupRoutes()
	r.Run(":8080")
}
