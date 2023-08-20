package router

import (
	notehandler "github.com/devghor/go-starter/internal/handler/note"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	note := api.Group("/note")
	note.Post("/", notehandler.CreateNotes)
	note.Get("/", notehandler.GetNotes)
	note.Get("/:noteId", notehandler.GetNote)
	note.Put("/:noteId", notehandler.UpdateNote)
	note.Delete("/:noteId", notehandler.DeleteNote)
}
