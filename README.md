# notes-api

A simple notes API built with Go.

## Project Structure

```bash
   notes-api/
   │── main.go          # Entry point for the application
   │── go.mod           # Go module file
   │── config/          # Configuration (DB setup, etc.)
   │── controllers/     # Request handlers (Controllers)
   │   │── note_controller.go
   │── services/        # Business logic (Service layer)
   │   │── note_service.go
   │── models/          # Data structures (Models)
   │   │── note.go
   │── routes/          # Routing setup
   │   │── routes.go
   │── storage/         # Data persistence (DB interactions)
   │   │── db.go
   │── utils/           # Helper functions

## Getting Started

### Prerequisites

- Go 1.20+ installed on your machine
- A running instance of a database (e.g., PostgreSQL)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/notes-api.git
   cd notes-api# Trigger Workflow
