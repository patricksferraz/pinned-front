package app

import (
	"embed"
	"fmt"
	"log"
	"net/http"

	"github.com/c-4u/pinned-front/config"
	"github.com/c-4u/pinned-front/domain/service"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html"
)

//go:embed views/*
var views embed.FS

//go:embed public/*
var public embed.FS

var sessionStore = session.New()

func init() {
	sessionStore.RegisterType(Session{})
}

func StartApp(conf *config.Config) {
	engine := html.NewFileSystem(http.FS(views), ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "views/layouts/main",
	})

	r := resty.New()
	service := service.NewService(r, conf.Baseurl)
	front := NewFront(service, sessionStore)
	middleware := NewMiddleware(sessionStore, service)

	app.Get("/", front.Index)
	guests := app.Group("/guests")
	{
		guests.Get("", middleware.CsrfProtection, front.SearchGuests)
		guests.Get("/:guest_id", front.GetGuest)
		guests.Get("/create", middleware.CsrfProtection, front.GetCreateGuest)
		guests.Post("/create", middleware.CsrfProtection, front.PostCreateGuest)
	}
	employees := app.Group("/employees")
	{
		employees.Get("", middleware.CsrfProtection, front.SearchEmployees)
		employees.Get("/:employee_id", front.GetEmployee)
		employees.Get("/create", middleware.CsrfProtection, front.GetCreateEmployee)
		employees.Post("/create", middleware.CsrfProtection, front.PostCreateEmployee)
	}
	places := app.Group("/places")
	{
		places.Get("", middleware.CsrfProtection, front.SearchPlaces)
		places.Get("/:place_id", front.GetPlace)
		places.Get("/create", middleware.CsrfProtection, front.GetCreatePlace)
		places.Post("/create", middleware.CsrfProtection, front.PostCreatePlace)
	}
	menus := app.Group("/menus")
	{
		menus.Get("", middleware.CsrfProtection, front.SearchMenus)
		menus.Get("/:menu_id", front.GetMenu)
		menus.Get("/create", middleware.CsrfProtection, front.GetCreateMenu)
		menus.Post("/create", middleware.CsrfProtection, front.PostCreateMenu)
	}

	app.Use(recover.New())

	app.Use("/static", filesystem.New(filesystem.Config{
		Root:       http.FS(public),
		PathPrefix: "public",
		Browse:     false,
	}))

	app.Use(func(c *fiber.Ctx) error {
		return c.Render("views/errors/error", fiber.Map{
			"Status": fmt.Sprintf("%d - %s", fiber.StatusNotFound, fiber.ErrNotFound),
			"Error":  c.Context().Request.URI()},
		) // => 404 "Not Found"
	})

	err := app.Listen("0.0.0.0:8080")
	if err != nil {
		log.Fatal("cannot start app", err)
	}
}
