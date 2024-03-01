package middleware

import (
	"github.com/gofiber/fiber/v2"
)

type guestMiddleware struct{}

func NewGuestMiddleware() Middleware {
	return guestMiddleware{}
}

func (m guestMiddleware) Execute() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("userID", float64(0))

		return c.Next()
	}
}
