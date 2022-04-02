package app

import (
	"fmt"

	"github.com/c-4u/pinned-front/domain/service"
	"github.com/c-4u/pinned-front/utils"
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
	var req CreateGuestSchema

	if err := c.BodyParser(&req); err != nil {
		return c.Render("views/errors/error", fiber.Map{
			"Status": fmt.Sprintf("%d - %s", fiber.StatusInternalServerError, fiber.ErrInternalServerError),
			"Error":  err.Error()},
		)
	}

	httpRes, httpError := a.Service.CreateGuest(c.Context(), req.Name, req.Doc)

	return c.Render("views/guests/create", fiber.Map{
		"Response":  httpRes,
		"Error":     httpError,
		"csrfToken": c.Locals("token"),
	})
}

func (a *Front) GetGuests(c *fiber.Ctx) error {
	var req SearchGuestsSchema

	if err := c.QueryParser(&req); err != nil {
		return c.Render("views/errors/error", fiber.Map{
			"Status": fmt.Sprintf("%d - %s", fiber.StatusInternalServerError, fiber.ErrInternalServerError),
			"Error":  err.Error()},
		)
	}

	if req.PageToken == nil {
		req.PageToken = utils.PString("")
	}

	if req.PageSize == nil {
		req.PageSize = utils.PInt(10)
	}

	httpRes, httpError := a.Service.SearchGuests(c.Context(), req.PageToken, req.PageSize)

	return c.Render("views/guests/index", fiber.Map{
		"Guests":        httpRes.Guests,
		"PageSize":      req.PageSize,
		"NextPageToken": httpRes.NextPageToken,
		"Error":         httpError,
	})
}

func (a *Front) GetCreateEmployee(c *fiber.Ctx) error {
	return c.Render("views/employees/create", fiber.Map{
		"csrfToken": c.Locals("token"),
	})
}

func (a *Front) PostCreateEmployee(c *fiber.Ctx) error {
	var req CreateEmployeeSchema

	if err := c.BodyParser(&req); err != nil {
		return c.Render("views/errors/error", fiber.Map{
			"Status": fmt.Sprintf("%d - %s", fiber.StatusInternalServerError, fiber.ErrInternalServerError),
			"Error":  err.Error()},
		)
	}

	httpRes, httpError := a.Service.CreateEmployee(c.Context(), req.Name)

	return c.Render("views/employees/create", fiber.Map{
		"Response":  httpRes,
		"Error":     httpError,
		"csrfToken": c.Locals("token"),
	})
}

func (a *Front) GetEmployees(c *fiber.Ctx) error {
	var req SearchEmployeesSchema

	if err := c.QueryParser(&req); err != nil {
		return c.Render("views/errors/error", fiber.Map{
			"Status": fmt.Sprintf("%d - %s", fiber.StatusInternalServerError, fiber.ErrInternalServerError),
			"Error":  err.Error()},
		)
	}

	if req.PageToken == nil {
		req.PageToken = utils.PString("")
	}

	if req.PageSize == nil {
		req.PageSize = utils.PInt(10)
	}

	httpRes, httpError := a.Service.SearchEmployees(c.Context(), req.PageToken, req.PageSize)

	return c.Render("views/employees/index", fiber.Map{
		"Employees":     httpRes.Employees,
		"PageSize":      req.PageSize,
		"NextPageToken": httpRes.NextPageToken,
		"Error":         httpError,
	})
}

func (a *Front) GetCreatePlace(c *fiber.Ctx) error {
	return c.Render("views/places/create", fiber.Map{
		"csrfToken": c.Locals("token"),
	})
}

func (a *Front) PostCreatePlace(c *fiber.Ctx) error {
	var req CreatePlaceSchema

	if err := c.BodyParser(&req); err != nil {
		return c.Render("views/errors/error", fiber.Map{
			"Status": fmt.Sprintf("%d - %s", fiber.StatusInternalServerError, fiber.ErrInternalServerError),
			"Error":  err.Error()},
		)
	}

	httpRes, httpError := a.Service.CreatePlace(c.Context(), req.Name)

	return c.Render("views/places/create", fiber.Map{
		"Response":  httpRes,
		"Error":     httpError,
		"csrfToken": c.Locals("token"),
	})
}

func (a *Front) GetPlaces(c *fiber.Ctx) error {
	var req SearchPlacesSchema

	if err := c.QueryParser(&req); err != nil {
		return c.Render("views/errors/error", fiber.Map{
			"Status": fmt.Sprintf("%d - %s", fiber.StatusInternalServerError, fiber.ErrInternalServerError),
			"Error":  err.Error()},
		)
	}

	if req.PageToken == nil {
		req.PageToken = utils.PString("")
	}

	if req.PageSize == nil {
		req.PageSize = utils.PInt(10)
	}

	httpRes, httpError := a.Service.SearchPlaces(c.Context(), req.PageToken, req.PageSize)

	return c.Render("views/places/index", fiber.Map{
		"Places":        httpRes.Places,
		"PageSize":      req.PageSize,
		"NextPageToken": httpRes.NextPageToken,
		"Error":         httpError,
	})
}

func (a *Front) GetCreateMenu(c *fiber.Ctx) error {
	return c.Render("views/menus/create", fiber.Map{
		"csrfToken": c.Locals("token"),
	})
}

func (a *Front) PostCreateMenu(c *fiber.Ctx) error {
	var req CreateMenuSchema

	if err := c.BodyParser(&req); err != nil {
		return c.Render("views/errors/error", fiber.Map{
			"Status": fmt.Sprintf("%d - %s", fiber.StatusInternalServerError, fiber.ErrInternalServerError),
			"Error":  err.Error()},
		)
	}

	httpRes, httpError := a.Service.CreateMenu(c.Context(), req.Name)

	return c.Render("views/menus/create", fiber.Map{
		"Response":  httpRes,
		"Error":     httpError,
		"csrfToken": c.Locals("token"),
	})
}

func (a *Front) GetMenus(c *fiber.Ctx) error {
	var req SearchMenusSchema

	if err := c.QueryParser(&req); err != nil {
		return c.Render("views/errors/error", fiber.Map{
			"Status": fmt.Sprintf("%d - %s", fiber.StatusInternalServerError, fiber.ErrInternalServerError),
			"Error":  err.Error()},
		)
	}

	if req.PageToken == nil {
		req.PageToken = utils.PString("")
	}

	if req.PageSize == nil {
		req.PageSize = utils.PInt(10)
	}

	httpRes, httpError := a.Service.SearchMenus(c.Context(), req.PageToken, req.PageSize)

	return c.Render("views/menus/index", fiber.Map{
		"Menus":         httpRes.Menus,
		"PageSize":      req.PageSize,
		"NextPageToken": httpRes.NextPageToken,
		"Error":         httpError,
	})
}
