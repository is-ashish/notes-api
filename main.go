package main

import (
	"notes-api/routes"
)

func main() {
	r := routes.SetupRoutes()
	r.Run(":8080")
}
