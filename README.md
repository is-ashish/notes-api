# notes-api

A simple notes API built with Go.

## Project Structure
`
 ```bash
    notes-api/
    │── .github/workflows/    # CI/CD pipeline for application
    │── config/               # Any app configs (e.g., env variables)
    │── controllers/          # API request handlers
    │   │── note_controller.go
    │── services/             # Business logic layer
    │   │── note_service.go
    │── models/               # Data models (e.g., Notes struct)
    │   │── note.go
    │── routes/               # API routing definitions
    │   │── routes.go
    │── storage/              # Database interaction layer
    │   │── db.go
    │── utils/                # Helper functions
    │── Dockerfile            # Docker configuration for app
    │── Makefile              # Automates build, test, deploy tasks
    │── go.mod                # Go module dependencies
    │── main.go               # App entry point
    │── README.md             # Documentation


## Getting Started

### Prerequisites

- Go 1.20+ installed on your machine
- A running instance of a database (e.g., PostgreSQL)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/notes-api.git
   cd notes-api


## Trigger Workflow
# notes-api/ CI/CD → Builds & pushes Docker image to Azure Container Registry (ACR) ✅ 
# notes-infra/ CI/CD → Deploys new Kubernetes configurations to Azure Kubernetes Service (AKS) ✅