package controllers

import (
	"net/http"
	"strconv"

	"notes-api/models"
	"notes-api/services"

	"github.com/gin-gonic/gin"
)

// Get all notes
func GetNotes(c *gin.Context) {
	c.JSON(http.StatusOK, services.GetNotes())
}

// Create a new note
func CreateNote(c *gin.Context) {
	var newNote models.Note
	if err := c.BindJSON(&newNote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, services.CreateNote(newNote))
}

// Update an existing note
func UpdateNote(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedNote models.Note
	if err := c.BindJSON(&updatedNote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if note, success := services.UpdateNote(id, updatedNote); success {
		c.JSON(http.StatusOK, note)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Note not found"})
	}
}

// Delete a note
func DeleteNote(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if success := services.DeleteNote(id); success {
		c.JSON(http.StatusOK, gin.H{"message": "Note deleted"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Note not found"})
	}
}
