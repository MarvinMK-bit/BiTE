package chesspuzzle

import (
	"github.com/Tibz-Dankan/BiTE/internal/models"
	"github.com/gofiber/fiber/v2"
)

var PostChessPuzzleAttempt = func(c *fiber.Ctx) error {
	chessPuzzleAttempt := models.ChessPuzzle{}

	response := map[string]interface{}{
		"status":  "success",
		"message": "Chess Puzzle Attempt created successfully!",
		"data":    chessPuzzleAttempt,
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}
