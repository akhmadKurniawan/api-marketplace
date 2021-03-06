//go:build wireinject
// +build wireinject

package main

import (
	"app/application/infrastructure/repository"
	"app/application/use_case/costumer/create_costumer"
	"app/application/use_case/login"
	"app/application/use_case/product/create_product"
	"app/application/use_case/product/get_product_shopid"
	"app/application/use_case/product_type/create_product_type"
	"app/application/use_case/seller/create_seller"
	"app/application/use_case/seller/delete_seller"
	"app/application/use_case/shop/create_shop"
	"app/application/use_case/transaction/create_transaction"
	"app/application/use_case/transaction/get_transaction"
	"app/application/use_case/transaction/scheduler_status"
	"app/application/use_case/transaction/update_transaction"
	"app/application/use_case/user/create_user"
	"app/application/use_case/user/delete_user"
	"app/application/use_case/user/update_user"
	"app/application/use_case/user/verify_email_user"
	"app/application/use_case/walet/create_walet"

	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func CreateUserHandler(db *gorm.DB) create_user.CreateUserHandler {
	wire.Build(create_user.NewCreateUserHandler, create_user.NewCreateUserService, repository.NewUserRepository, repository.NewSellerRepository, repository.NewCostumerRepository)
	return create_user.CreateUserHandler{}
}

func UpdateUserHandler(db *gorm.DB) update_user.UpdateUserHandler {
	wire.Build(update_user.NewUpdateUserHandler, update_user.NewUpdateUserService, repository.NewUserRepository)
	return update_user.UpdateUserHandler{}
}

func DeleteUserHandler(db *gorm.DB) delete_user.DeleteUserHandler {
	wire.Build(delete_user.NewDeleteUserHandler, delete_user.NewDeleteUserService, repository.NewUserRepository)
	return delete_user.DeleteUserHandler{}
}

func VerifyEmailUserHandler(db *gorm.DB) verify_email_user.VerifyEmailUserHandler {
	wire.Build(verify_email_user.NewVerifyEmailUserHandler, verify_email_user.NewVerifyEmailUserService, repository.NewUserRepository)
	return verify_email_user.VerifyEmailUserHandler{}
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
	wire.Build(create_costumer.NewCreateCostumerHandler, create_costumer.NewCreateCostumerService, repository.NewCostumerRepository)
	return create_costumer.CreateCostumerHandler{}
}

func CreateProductHandler(db *gorm.DB, mongo *mongo.Database) create_product.CreateProductHandler {
	wire.Build(create_product.NewCreateProductHandler, create_product.NewCreateProductService, repository.NewProductRepository, repository.NewShopRepository, repository.NewProductTypeRepository)
	return create_product.CreateProductHandler{}
}

func ShowProductByShopIDHandler(db *gorm.DB, mongo *mongo.Database) get_product_shopid.ShowProductByShopIDHandler {
	wire.Build(get_product_shopid.NewShowProductByShopIDHandler, get_product_shopid.NewShowProductByShopIDService, repository.NewProductRepository)
	return get_product_shopid.ShowProductByShopIDHandler{}
}

func CreateProductTypeHandler(db *gorm.DB, mongo *mongo.Database) create_product_type.CreateProductTypeHandler {
	wire.Build(create_product_type.NewCreateProductTypeHandler, create_product_type.NewCreateProductTypeService, repository.NewProductTypeRepository)
	return create_product_type.CreateProductTypeHandler{}
}

func CreateShopHandler(db *gorm.DB) create_shop.CreateShopHandler {
	wire.Build(create_shop.NewCreateShopHandler, create_shop.NewCreateShopService, repository.NewShopRepository, repository.NewSellerRepository)
	return create_shop.CreateShopHandler{}
}

func CreateTransactionHandler(db *gorm.DB, mongo *mongo.Database) create_transaction.CreateTransactionHandler {
	wire.Build(create_transaction.NewCreateTransactionHandler, create_transaction.NewCreateTransactionService, repository.NewTransactionRepository, repository.NewProductRepository, repository.NewWaletRepository, repository.NewSellerRepository, repository.NewShopRepository)
	return create_transaction.CreateTransactionHandler{}
}

func ShowTransactionHandler(db *gorm.DB) get_transaction.ShowTransactionHandler {
	wire.Build(get_transaction.NewShowTransactionHandler, get_transaction.NewShowTransactionService, repository.NewTransactionRepository)
	return get_transaction.ShowTransactionHandler{}
}

func UpdateTransactionHandler(db *gorm.DB) update_transaction.UpdateTransactionHandler {
	wire.Build(update_transaction.NewUpdateTransactionHandler, update_transaction.NewUpdateTransactionService, repository.NewTransactionRepository)
	return update_transaction.UpdateTransactionHandler{}
}

func CreateWaletHandler(db *gorm.DB) create_walet.CreateWaletHandler {
	wire.Build(create_walet.NewCreateWaletHandler, create_walet.NewCreateWaletService, repository.NewWaletRepository)
	return create_walet.CreateWaletHandler{}
}

func UpdateSchedulerHandler(db *gorm.DB) scheduler_status.UpdateSchedulerHandler {
	wire.Build(scheduler_status.NewUpdateSchedulerHandler, scheduler_status.NewUpdateSchedulerService, repository.NewTransactionRepository)
	return scheduler_status.UpdateSchedulerHandler{}
}
