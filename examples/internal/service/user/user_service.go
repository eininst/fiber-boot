package user

import (
	"context"
	"github.com/eininst/flog"
)

type UserService interface {
	Add(ctx context.Context)
}
type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}

func (us *userService) Add(ctx context.Context) {
	flog.Info("Add a user", ctx)
}
