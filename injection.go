//go:build wireinject
// +build wireinject

package main

import (
	"app/application/use_case/user/create_user"
	"app/repository"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func CreateUserHandler(db *gorm.DB) create_user.CreateUserHandler {
	wire.Build(create_user.NewCreateUserHandler, create_user.NewCreateUserService, repository.NewUserRepository)
	return create_user.CreateUserHandler{}
}
