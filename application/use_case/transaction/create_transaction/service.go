package create_transaction

import (
	"app/application/infrastructure"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/virtualaccount"
)

type CreateTransactionService struct {
	transactionRepository infrastructure.TransactionRepository
	productRepository     infrastructure.ProductRepository
	waletRepository       infrastructure.WaletRepository
	sellerRepository      infrastructure.SellerRepository
	shopRepository        infrastructure.ShopRepository
}

func NewCreateTransactionService(transactionRepo infrastructure.TransactionRepository, productRepo infrastructure.ProductRepository, waletRepo infrastructure.WaletRepository, sellerRepo infrastructure.SellerRepository, shopRepo infrastructure.ShopRepository) CreateTransactionService {
	return CreateTransactionService{
		transactionRepository: transactionRepo,
		productRepository:     productRepo,
		waletRepository:       waletRepo,
		sellerRepository:      sellerRepo,
		shopRepository:        shopRepo,
	}
}

func (s *CreateTransactionService) CreateTransaction(ctx context.Context, req CreateTransactionRequest) (string, error) {
	product, errProduct := s.productRepository.GetProductByID(ctx, req.ProductID)
	if errProduct != nil {
		log.Println("Service - CreateTransaction errorProduct : ", errProduct)
		return "", errProduct
	}
	if product.Qty < req.TotalProduct {
		errProduct = errors.New("out of stok")
		return "", errProduct
	}
	qty := product.Qty - req.TotalProduct
	product.Qty = qty
	_, errUp := s.productRepository.UpdateProdut(ctx, product.ID, product.Qty)
	if errUp != nil {
		log.Println("Service - CreateTransaction errorUpdateProduct : ", errUp)
		return "", errUp
	}

	// get id buyer and update the saldo
	walet, errWalet := s.waletRepository.GetWaletByUserID(ctx, req.UserID)
	if errWalet != nil {
		log.Println("Service - CreateTransaction errorWalet : ", errWalet)
		return "", errWalet
	}

	totalAmount := product.Price * req.TotalProduct
	if walet.Saldo < totalAmount {
		errWalet = errors.New("your saldo not enough")
		return "", errWalet
	}

	xendit.Opt.SecretKey = os.Getenv("KEY_XENDIT")
	ut := virtualaccount.CreateFixedVAParams{IsSingleUse: func() *bool { b := true; return &b }()}
	ct := virtualaccount.CreateFixedVAParams{IsClosed: func(b bool) *bool { return &b }(true)}
	IdVa := "VA_fixed-" + time.Now().Format("2006010205")
	data := virtualaccount.CreateFixedVAParams{
		ExternalID:     IdVa,
		BankCode:       "BRI",
		Name:           "Akhmad Kurniawan",
		ExpectedAmount: float64(totalAmount),
		IsClosed:       *&ct.IsClosed,
		IsSingleUse:    *&ut.IsSingleUse,
	}

	resp, err := virtualaccount.CreateFixedVA(&data)
	if err != nil {
		log.Println("Service - CreateTransaction errXendit : ", err)
		return "", err
	}

	fmt.Printf("created fixed va: %+v\n", resp)

	saldo := walet.Saldo - totalAmount
	walet.Saldo = saldo
	_, errUpWalet := s.waletRepository.UpdateWaletSaldo(ctx, req.UserID, walet.Saldo)
	if errUpWalet != nil {
		log.Println("Service - CreateTransaction errorUpWalet : ", errUpWalet)
		return "", errUpWalet
	} //get id buyer and update the saldo

	// get id seller and update the saldo
	shop, errShop := s.shopRepository.GetShopBySellerId(ctx, product.ShopId)
	if errShop != nil {
		log.Println("Service - CreateTransaction errorShop : ", errShop)
		return "", errShop
	}

	seller, errSeller := s.sellerRepository.GetSellerByID(ctx, shop.SellerID)
	if errSeller != nil {
		log.Println("Service - CreateTransaction errorSeller : ", errSeller)
		return "", errSeller
	}

	waletSeller, errSellerWalet := s.waletRepository.GetWaletByUserID(ctx, seller.UserID)
	if errSellerWalet != nil {
		log.Println("Service - CreateTransaction errorSellerWalet : ", errSellerWalet)
		return "", errSellerWalet
	}

	saldoSeller := waletSeller.Saldo + totalAmount
	walet.Saldo = saldoSeller
	_, errUpWalet = s.waletRepository.UpdateWaletSaldo(ctx, seller.UserID, walet.Saldo)
	if errUpWalet != nil {
		log.Println("Service - CreateTransaction errorUpWalet : ", errUpWalet)
		return "", errUpWalet
	} // get id seller and update the saldo

	status := "pending"
	debit := totalAmount - (totalAmount + totalAmount)
	req.ProductID = product.ID
	tp := req.TotalProduct - (req.TotalProduct + req.TotalProduct)
	errCreate := s.transactionRepository.CreateTransaction(ctx, RequestMapper(req, debit, "Debit", status, resp.ExternalID, tp))
	if errCreate != nil {
		log.Println("Service - CreateTransaction errorCreate : ", errCreate)
		return "", errCreate
	}

	errCreate = s.transactionRepository.CreateTransaction(ctx, RequestMapperK(req, totalAmount, "Kredit", status, resp.ExternalID, seller.UserID))
	if errCreate != nil {
		log.Println("Service - CreateTransaction errorCreate : ", errCreate)
		return "", errCreate
	}

	return resp.ExternalID, nil
}
