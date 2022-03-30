package app

import (
	"time"

	"github.com/c-4u/pinned-front/domain/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
)

type Middleware struct {
	Session        *session.Store
	Service        *service.Service
	CsrfProtection func(*fiber.Ctx) error
}

func NewMiddleware(session *session.Store, service *service.Service) *Middleware {
	m := &Middleware{
		Session: session,
		Service: service,
	}
	m.csrfProtection()
	return m
}

func (m *Middleware) csrfProtection() {
	m.CsrfProtection = csrf.New(csrf.Config{
		KeyLookup:      "form:_csrf",
		CookieName:     "csrf_",
		CookieSameSite: "Strict",
		Expiration:     1 * time.Hour,
		KeyGenerator:   utils.UUID,
		ContextKey:     "token",
	})
}

// func (m *Middleware) InitSession(c *fiber.Ctx) error {
// 	currSession, err := m.Session.Get(c)
// 	if err != nil {
// 		return err
// 	}

// 	r := currSession.Get("data")
// 	if r == nil {
// 		currSession.Set("data", Session{})
// 		currSession.Save()
// 	}

// 	return c.Next()
// }
