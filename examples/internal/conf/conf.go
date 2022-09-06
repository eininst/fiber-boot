package conf

import (
	fboot "github.com/eininst/fiber-boot"
	"github.com/eininst/fiber-boot/examples/internal/service/user"
)

func Provide() {
	//inject services
	fboot.Provide(user.NewUserService())
}
