package main

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
)

type Comment struct {
	// Text is a field that represents a text input. It is expected to be provided in form data and JSON payloads.
	// The "binding:\"required\"" tag indicates that this field is mandatory and must be provided in the request.
	Text string `form:"text" json:"text" binding:"required"`
	User string `form:"user" json:"user" binding:"required"`
}

func main() {
	app := fiber.New()
	api := app.Group("/api/v1")
	api.Post("/comment", createComment)
	app.Listen(":3000")
}

func createComment(c fiber.Ctx) error {
	Comment := new(Comment)
	err := json.Unmarshal(c.Body(), Comment)
	// text in the commnt is required field

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
			"error":   err.Error(),
		})
	}

	PushCommentToQueue(Comment)

	// return a response
	return c.Status(200).JSON(fiber.Map{
		"message": "Comment created successfully",
		"comment": Comment.Text,
	})
}
