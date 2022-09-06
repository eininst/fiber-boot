package fboot

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

type App struct {
	*fiber.App
}

type Api interface {
	Router(r fiber.Router)
}

func New(config ...fiber.Config) *App {
	return &App{
		App: fiber.New(config...),
	}
}

func (a *App) Install(api Api) {
	//Injector
	Provide(api)
	Populate()
	api.Router(a.App)
}

func (a *App) Listen(addr string) {
	Listen(a.App, addr)
}

func (a *App) ListenTLS(addr, certFile, keyFile string, timeout ...time.Duration) {
	ListenTLS(a.App, addr, certFile, keyFile, timeout...)
}
