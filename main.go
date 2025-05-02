package main

import (
	"notes-api/routes"
)

func main() {
	r := routes.SetupRoutes()
	r.Run(":8080")
}

// notes-api/
// │── main.go          # Entry point for the application
// │── go.mod           # Go module file
// │── config/          # Configuration (DB setup, etc.)
// │── controllers/     # Request handlers (Controllers)
// │   │── note_controller.go
// │── services/        # Business logic (Service layer)
// │   │── note_service.go
// │── models/          # Data structures (Models)
// │   │── note.go
// │── routes/          # Routing setup
// │   │── routes.go
// │── storage/         # Data persistence (DB interactions)
// │   │── db.go
// │── utils/           # Helper functions
