//go:build wireinject
// +build wireinject

package main

import (
	"app/application/use_case/costumer/create_costumer"
	"app/application/use_case/login"
	"app/application/use_case/product/create_product"
	"app/application/use_case/product_type/create_product_type"
	"app/application/use_case/seller/create_seller"
	"app/application/use_case/seller/delete_seller"
	"app/application/use_case/user/create_user"
	"app/application/use_case/user/delete_user"
	"app/repository"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func CreateUserHandler(db *gorm.DB) create_user.CreateUserHandler {
	wire.Build(create_user.NewCreateUserHandler, create_user.NewCreateUserService, repository.NewUserRepository)
	return create_user.CreateUserHandler{}
}

func DeleteUserHandler(db *gorm.DB) delete_user.DeleteUserHandler {
	wire.Build(delete_user.NewDeleteUserHandler, delete_user.NewDeleteUserService, repository.NewUserRepository)
	return delete_user.DeleteUserHandler{}
}

func LoginHandler(db *gorm.DB) login.LoginHandler {
	wire.Build(login.NewLoginHandler, login.NewLoginService, repository.NewLoginRepository, repository.NewUserRepository)
	return login.LoginHandler{}
}

func CreateSellerHandler(db *gorm.DB) create_seller.CreateSellerHandler {
	wire.Build(create_seller.NewCreateSellerHandler, create_seller.NewCreateSellerService, repository.NewSellerRepository)
	return create_seller.CreateSellerHandler{}
}

func DeleteSellerHandler(db *gorm.DB) delete_seller.DeleteSellerHandler {
	wire.Build(delete_seller.NewDeleteSellerHandler, delete_seller.NewDeleteSellerService, repository.NewSellerRepository)
	return delete_seller.DeleteSellerHandler{}
}

func CreateCostumerHandler(db *gorm.DB) create_costumer.CreateCostumerHandler {
	wire.Build(create_costumer.NewCreateCostumerHandler, create_costumer.NewCreateCostumerService, repository.NewCostumerRepository, repository.NewUserRepository)
	return create_costumer.CreateCostumerHandler{}
}

func CreateProductHandler(db *gorm.DB) create_product.CreateProductHandler {
	wire.Build(create_product.NewCreateProductHandler, create_product.NewCreateProductService, repository.NewProductRepository)
	return create_product.CreateProductHandler{}
}

func CreateProductTypeHandler(db *gorm.DB) create_product_type.CreateProductTypeHandler {
	wire.Build(create_product_type.NewCreateProductTypeHandler, create_product_type.NewCreateProductTypeService, repository.NewProductTypeRepository)
	return create_product_type.CreateProductTypeHandler{}
}
