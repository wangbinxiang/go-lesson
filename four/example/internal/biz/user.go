package biz

import "context"

type User struct {
	Id   int
	Name string
	City string
}

type UserUsecase interface {
	GetUserInfo(ctx context.Context, id int) (*User, error)
}

type UserRepo interface {
	GetUserInfo(ctx context.Context, id int) (*User, error)
}
