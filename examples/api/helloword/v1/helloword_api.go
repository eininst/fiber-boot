package v1

import (
	fboot "github.com/eininst/fiber-boot"
	"github.com/eininst/fiber-boot/examples/internal/service/user"
	"github.com/gofiber/fiber/v2"
)

func init() {
	fboot.Provide(new(HellowordApi))
}

type HellowordApi struct {
	UserService user.UserService `inject:""`
}

// @Summary 测试swagger
// @Tags test
// @version 1.0

// @Router / [get]
func (h *HellowordApi) Add(c *fiber.Ctx) error {
	h.UserService.Add(c.Context())
	return c.JSON("hello123")
}
