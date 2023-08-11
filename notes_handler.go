package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (dataSource *DataSource) getAllNotes(c *fiber.Ctx) error {
	var res string
	var todos []string

	rows, err := dataSource.Query("SELECT id FROM notes")
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{"message": err})
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&res)
		todos = append(todos, res)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"notes": todos})
}

func (dataSource *DataSource) addNote(c *fiber.Ctx) error {
	_, err := dataSource.Exec("INSERT into notes VALUES ($1, $2, $3, $4)", uuid.New(), time.Now(), time.Now(), "First Note")
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{"message": err})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Note Created Successfully"})
}
