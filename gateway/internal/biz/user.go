package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Hello string
}

type UserRepo interface {
	CreateUser(*User) error
	UpdateUser(*User) error
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper("usecase/user", logger)}
}

func (uc *UserUsecase) Create(g *User) error {
	return uc.repo.CreateUser(g)
}

func (uc *UserUsecase) Update(g *User) error {
	return uc.repo.UpdateUser(g)
}
