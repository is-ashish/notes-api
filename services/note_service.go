package services

import "notes-api/models"

// Dummy data
var notes = []models.Note{
	{ID: 1, Title: "First Note", Content: "This is the first note"},
	{ID: 2, Title: "Second Note", Content: "This is the second note"},
}

func GetNotes() []models.Note {
	return notes
}

func CreateNote(note models.Note) models.Note {
	note.ID = len(notes) + 1
	notes = append(notes, note)
	return note
}

func UpdateNote(id int, updatedNote models.Note) (models.Note, bool) {
	for i, note := range notes {
		if note.ID == id {
			notes[i] = updatedNote
			return updatedNote, true
		}
	}
	return models.Note{}, false
}

func DeleteNote(id int) bool {
	for i, note := range notes {
		if note.ID == id {
			notes = append(notes[:i], notes[i+1:]...)
			return true
		}
	}
	return false
}
