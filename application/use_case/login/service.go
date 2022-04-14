package login

import (
	"app/application/infrastructure"
	"app/middleware"
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	loginRepository infrastructure.LoginRepository
	userRepository  infrastructure.UserRepository
}

func NewLoginService(loginRepo infrastructure.LoginRepository, userRepo infrastructure.UserRepository) LoginService {
	return LoginService{
		loginRepository: loginRepo,
		userRepository:  userRepo,
	}
}

func (s *LoginService) LoginUser(ctx context.Context, req LoginRequest) (*Response, error) {
	// get username
	user, errUser := s.userRepository.GetAllUsername(ctx, req.Username)
	if errUser != nil {
		log.Debug().
			Str("Func", "LoginUser").
			Str("Usernam", "LoginUser").
			Interface("Request", req).
			Msg("Service - Login error while access username")
		errUser = errors.New("Service - Login error while access username")
		return nil, errUser
	}
	byteDBPass := []byte(user.Password)
	byteReqPass := []byte(req.Password)

	//compare hash password
	if error := bcrypt.CompareHashAndPassword(byteDBPass, byteReqPass); error != nil {
		log.Error().Err(errUser).Msg("Service - Compare hash error")
		return nil, error
	}

	//create Claims
	claims := middleware.CreateClaims(uint64(user.ID), user.Username, user.Role, time.Duration(64))
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims) // create token
	signed, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		log.Error().Err(err).Msg("Service - Signed Token error")
		return nil, err
	}
	tes, errLogin := s.loginRepository.Login(ctx, RequestMapper(user.ID, signed), user.ID)
	if errLogin != nil {
		log.Error().Err(errLogin).Msg("Service - Login error")
		return nil, errLogin
	}
	fmt.Println(tes)

	userId := strconv.FormatUint(uint64(user.ID), 10)

	//Get user By Id
	userData, errGetUser := s.userRepository.GetUserID(ctx, userId)
	if errGetUser != nil {
		log.Debug().
			Str("Func", "LoginUser").
			Str("UserID", "LoginUser").
			Interface("Request", req).
			Msg("Service - GetUserID error")
		errGetUser = errors.New("Service - GetUserID error")
		return nil, errGetUser
	}
	// fmt.Println("atas")
	// if errGetUser != nil || userData != nil || userData.Status == "Inactivated" {
	// 	log.Error().Err(errUser).Msg("please activate your account")
	// 	return nil, errUser
	// }
	// fmt.Println("bawah")

	return &Response{User: *userData}, nil

}
