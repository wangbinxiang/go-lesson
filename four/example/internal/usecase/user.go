package usecase

import (
	"context"

	"github.com/wangbinxiang/go-lesson-four/internal/biz"
)

type user struct {
	repo biz.UserRepo
}

func (u *user) GetUserInfo(ctx context.Context, id int) (*biz.User, error) {
	return u.repo.GetUserInfo(ctx, id)
}

func NewUserUsecase(repo biz.UserRepo) biz.UserUsecase {
	return &user{repo: repo}
}
