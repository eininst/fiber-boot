package main

import (
	fboot "github.com/eininst/fiber-boot"
	"github.com/eininst/fiber-boot/examples/api/helloword"
	"github.com/eininst/fiber-boot/examples/internal/conf"
	"github.com/gofiber/fiber/v2"
	"time"
)

func init() {
	fboot.SetConfig("./examples/configs/helloword.yml")
	conf.Provide()
}

func main() {
	r := fboot.New(fiber.Config{
		Prefork:     true,
		ReadTimeout: time.Second * 10,
	})
	r.Install(&helloword.Api{})

	port := fboot.Get("port").String()
	r.Listen(port)
}
