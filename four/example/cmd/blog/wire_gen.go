// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/wangbinxiang/go-lesson-four/internal/data"
	"github.com/wangbinxiang/go-lesson-four/internal/service"
	"github.com/wangbinxiang/go-lesson-four/internal/usecase"
)

// Injectors from wire.go:

func InitUserService() (*service.UserService, error) {
	userRepo := data.NewUserRepository()
	userUsecase := usecase.NewUserUsecase(userRepo)
	userService := service.NewUserService(userUsecase)
	return userService, nil
}