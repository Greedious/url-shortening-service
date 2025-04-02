package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	"url-shortner/configs"
	"url-shortner/routes"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/limiter"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/joho/godotenv"
)

var envConfig *configs.Config

func main() {
	loadEnvErr := godotenv.Load(".env")
	if loadEnvErr != nil {
		fmt.Println(loadEnvErr)
		return
	}

	envConfig = configs.LoadConfig()

	app := fiber.New()

	registerLogger(app)
	registerRateLimiter(app)

	setupRoutes(app)

	log.Fatal(app.Listen(":" + strconv.Itoa(envConfig.APIPort)))
}

func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api", routes.ShortenURL)
}

func registerLogger(app *fiber.App) {
	app.Use(logger.New())
}

func registerRateLimiter(app *fiber.App) {
	rateLimitingThreshold := os.Getenv("API_RATE_LIMIT_THRESHOLD")
	maxReqsPerMin, err := strconv.Atoi(rateLimitingThreshold)
	if err != nil {
		// Fallback to default 10
		maxReqsPerMin = 10
	}
	app.Use(limiter.New(limiter.Config{
		Max:        maxReqsPerMin,
		Expiration: time.Minute,
		LimitReached: func(ctx fiber.Ctx) error {
			return ctx.Status(fiber.StatusTooManyRequests).
				JSON(fiber.Map{"error": "Too many requests."})
		},
		SkipFailedRequests:     false,
		SkipSuccessfulRequests: false,
		LimiterMiddleware:      limiter.FixedWindow{},
	}))
}
