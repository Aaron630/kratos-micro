package data

import (
	"gateway/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func (r *userRepo) CreateUser(user *biz.User) error {
	return nil
}

func (r *userRepo) UpdateUser(user *biz.User) error {
	return nil
}

// NewGreeterRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper("data/user", logger),
	}
}

func (r *userRepo) CreateGreeter(g *biz.User) error {
	return nil
}

func (r *userRepo) UpdateGreeter(g *biz.User) error {
	return nil
}
