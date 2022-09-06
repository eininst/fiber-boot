package fboot

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strings"
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
	if !strings.HasPrefix(addr, ":") {
		addr = fmt.Sprintf(":%s", addr)
	}
	Listen(a.App, addr)
}

func (a *App) ListenTLS(addr, certFile, keyFile string, timeout ...time.Duration) {
	if !strings.HasPrefix(addr, ":") {
		addr = fmt.Sprintf(":%s", addr)
	}
	ListenTLS(a.App, addr, certFile, keyFile, timeout...)
}
