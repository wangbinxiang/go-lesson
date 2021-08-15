package data

import (
	"context"

	"github.com/wangbinxiang/go-lesson-four/internal/biz"
)

type userRepo struct {
}

func (u *userRepo) GetUserInfo(ctx context.Context, id int) (*biz.User, error) {
	return &biz.User{
		Id:   1,
		Name: "大明老师",
		City: "上海",
	}, nil
}

func NewUserRepository() biz.UserRepo {
	return &userRepo{}
}
