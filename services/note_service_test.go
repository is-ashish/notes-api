package services

import (
	"notes-api/models"
	"testing"
)

func TestCreateNote(t *testing.T) {
	note := CreateNote(models.Note{Title: "Test", Content: "Testing"})
	if note.Title != "Test" {
		t.Errorf("Expected Title 'Test', got %s", note.Title)
	}
}
