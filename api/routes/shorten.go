package routes

import (
	"fmt"
	"os"
	"time"
	"url-shortner/database"
	"url-shortner/helpers"
	"url-shortner/validation"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type request struct {
	URL    string        `json:"url" validate:"required"`
	Expiry time.Duration `json:"expiry"`
}

type response struct {
	URL         string        `json:"url" validate:"required"`
	Short       string        `json:"short"`
	Expiry      time.Duration `json:"expiry" validate:"required"`
	AccessCount int           `json: "access_count"`
}

func ShortenURL(ctx fiber.Ctx) error {
	body := new(request)
	domain := os.Getenv("API_DOMAIN")

	if err := ctx.Bind().Body(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse json"})
	}

	if validationErr := validation.ValidateStruct(body); validationErr != nil {
		return fmt.Errorf("%s", validationErr.Error())
	}

	// check if the given url in the body is in URL shape
	if !govalidator.IsURL(body.URL) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "the provided url property is not a URL"})
	}

	// making sure the url in the body is not same as our system URL
	if !helpers.RemoveDomainErrors(body.URL) {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Not allowed URL"})
	}

	// enforcing http
	body.URL = helpers.EnforceHTTP(body.URL)

	// Here we will assume that the generated uuid is unique, but in real world systems, it needs to be validated
	// TODO: Implement mechanism to make sure the new generated id is unique, for now we will just throw internal server error if it is found in DB
	id := uuid.NewString()[:7]

	if val, err := database.GetOriginalURL(id); err == nil || val != "" {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something wrong happened"})
	}

	body.Expiry = 24

	if err := database.StoreShortenedURL(id, body.URL, 24*3600*time.Second); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error happened in server"})
	}

	resp := response{
		URL:         body.URL,
		Short:       domain + "/" + id,
		Expiry:      body.Expiry,
		AccessCount: 0,
	}
	return ctx.Status(fiber.StatusOK).JSON(resp)
}
