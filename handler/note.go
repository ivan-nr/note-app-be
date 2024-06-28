package handler

import (
	"api-fiber-gorm/database"
	"api-fiber-gorm/model"

	"github.com/gofiber/fiber/v2"
)

func GetAllNotes(c *fiber.Ctx) error {
	db := database.DB
	var note []model.Note
	db.Find(&note)
	return c.JSON(fiber.Map{"status": "success", "message": "All Notes", "data": note})
}

func GetNote(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var note model.Note
	db.Find(&note, id)
	if note.Title == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note found with ID", "data": nil})

	}
	return c.JSON(fiber.Map{"status": "success", "message": "note found", "data": note})
}

func CreateNote(c *fiber.Ctx) error {
	db := database.DB
	note := new(model.Note)

	userId := c.Locals("user_id").(int)

	if err := c.BodyParser(note); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create note", "data": err})
	}

	note.UserId = userId

	db.Create(&note)

	return c.JSON(fiber.Map{"status": "success", "message": "Created note", "data": note})
}

func DeleteNote(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var note model.Note
	db.First(&note, id)
	if note.Title == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note found with ID", "data": nil})

	}
	db.Delete(&note)
	return c.JSON(fiber.Map{"status": "success", "message": "note successfully deleted", "data": nil})
}

func UpdateNote(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var note model.Note
	db.First(&note, id)
	if note.Title == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note found with ID", "data": nil})
	}
	note.Title = c.FormValue("title")
	note.Note = c.FormValue("note")
	db.Save(&note)
	return c.JSON(fiber.Map{"status": "success", "message": "note successfully updated", "data": note})
}
