//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/wangbinxiang/go-lesson-four/internal/data"
	"github.com/wangbinxiang/go-lesson-four/internal/service"
	"github.com/wangbinxiang/go-lesson-four/internal/usecase"
)

func InitUserService() (*service.UserService, error) {
	wire.Build(service.NewUserService, usecase.NewUserUsecase, data.NewUserRepository)
	return new(service.UserService), nil
}
