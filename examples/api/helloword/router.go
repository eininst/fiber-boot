package helloword

import (
	v1 "github.com/eininst/fiber-boot/examples/api/helloword/v1"
	"github.com/gofiber/fiber/v2"
)

type Api struct {
	Helloword *v1.HellowordApi `inject:""`
}

func (api *Api) Router(r fiber.Router) {
	r.Get("/add", api.Helloword.Add)
}
