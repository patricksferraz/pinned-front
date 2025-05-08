package app

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/patricksferraz/pinned-front/domain/service"
	"github.com/patricksferraz/pinned-front/utils"
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

func (a *Front) SearchGuests(c *fiber.Ctx) error {
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

func (a *Front) GetGuest(c *fiber.Ctx) error {
	guestID := c.Params("guest_id")
	if !govalidator.IsUUIDv4(guestID) {
		return c.Render("views/errors/error", fiber.Map{
			"Status": fmt.Sprintf("%d - %s", fiber.StatusBadRequest, fiber.ErrBadRequest),
			"Error":  "guest_id is not a valid uuid"},
		)
	}

	guest, httpError := a.Service.GetGuest(c.Context(), &guestID)
	return c.Render("views/guests/profile", fiber.Map{
		"Guest": guest,
		"Error": httpError,
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

func (a *Front) SearchEmployees(c *fiber.Ctx) error {
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

func (a *Front) GetEmployee(c *fiber.Ctx) error {
	employeeID := c.Params("employee_id")
	if !govalidator.IsUUIDv4(employeeID) {
		return c.Render("views/errors/error", fiber.Map{
			"Status": fmt.Sprintf("%d - %s", fiber.StatusBadRequest, fiber.ErrBadRequest),
			"Error":  "employee_id is not a valid uuid"},
		)
	}

	employee, httpError := a.Service.GetEmployee(c.Context(), &employeeID)
	return c.Render("views/employees/profile", fiber.Map{
		"Employee": employee,
		"Error":    httpError,
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

func (a *Front) SearchPlaces(c *fiber.Ctx) error {
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

func (a *Front) GetPlace(c *fiber.Ctx) error {
	placeID := c.Params("place_id")
	if !govalidator.IsUUIDv4(placeID) {
		return c.Render("views/errors/error", fiber.Map{
			"Status": fmt.Sprintf("%d - %s", fiber.StatusBadRequest, fiber.ErrBadRequest),
			"Error":  "place_id is not a valid uuid"},
		)
	}

	place, httpError := a.Service.GetPlace(c.Context(), &placeID)
	return c.Render("views/places/profile", fiber.Map{
		"Place": place,
		"Error": httpError,
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

func (a *Front) SearchMenus(c *fiber.Ctx) error {
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

func (a *Front) GetMenu(c *fiber.Ctx) error {
	menuID := c.Params("menu_id")
	if !govalidator.IsUUIDv4(menuID) {
		return c.Render("views/errors/error", fiber.Map{
			"Status": fmt.Sprintf("%d - %s", fiber.StatusBadRequest, fiber.ErrBadRequest),
			"Error":  "menu_id is not a valid uuid"},
		)
	}

	menu, httpError := a.Service.GetMenu(c.Context(), &menuID)
	return c.Render("views/menus/profile", fiber.Map{
		"Menu":  menu,
		"Error": httpError,
	})
}
