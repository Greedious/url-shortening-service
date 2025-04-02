package routes

import (
	"url-shortner/database"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v3"
)

func ResolveURL(ctx fiber.Ctx) error {
	url := ctx.Params("url")

	result, err := database.GetOriginalURL(url)

	if err == redis.Nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "short for the URL was not found"})
	} else if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot connect to Database"})
	}

	rInr := database.CreateClient(1)
	defer rInr.Close()

	// increasing the counter of total resolve calls
	_ = rInr.Incr(database.Ctx, "counter")

	return ctx.Redirect().Status(fiber.StatusMovedPermanently).To(result)
}
