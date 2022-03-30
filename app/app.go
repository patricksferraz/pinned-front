package app

import (
	"fmt"

	"github.com/c-4u/pinned-front/domain/entity"
	"github.com/c-4u/pinned-front/domain/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type Front struct {
	Service *service.Service
	Session *session.Store
}

func NewFront(service *service.Service, session *session.Store) *Front {
	return &Front{
		Service: service,
		Session: session,
	}
}

func (a *Front) Index(c *fiber.Ctx) error {
	return c.Render("views/index", fiber.Map{})
}

func (a *Front) GetCreateGuest(c *fiber.Ctx) error {
	return c.Render("views/guests/create", fiber.Map{
		"csrfToken": c.Locals("token"),
	})
}

func (a *Front) PostCreateGuest(c *fiber.Ctx) error {
	var req entity.CreateGuestRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Render("views/errors/error", fiber.Map{
			"Status": fmt.Sprintf("%d - %s", fiber.StatusInternalServerError, fiber.ErrInternalServerError),
			"Error":  err.Error()},
		)
	}

	httpRes, httpError := a.Service.CreateGuest(c.Context(), &req)

	return c.Render("views/guests/create", fiber.Map{
		"Response":  httpRes,
		"Error":     httpError,
		"csrfToken": c.Locals("token"),
	})
}
