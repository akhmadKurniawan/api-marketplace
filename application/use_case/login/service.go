package login

import (
	"app/application/infrastructure"
	"app/middleware"
	"context"
	"log"
	"strconv"
	"time"

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
	user, errUser := s.userRepository.GetUsername(ctx, req.Username)
	if errUser != nil {
		log.Println("Service - Login error while access username : ", errUser)
	}

	byteDBPass := []byte(user.Password)
	byteReqPass := []byte(req.Password)

	//compare hash password
	if error := bcrypt.CompareHashAndPassword(byteDBPass, byteReqPass); error != nil {
		log.Println("Service - Compare hash error : ", error)
	}

	//create Claims
	claims := middleware.CreateClaims(uint64(user.ID), user.Username, user.Role, time.Duration(64))
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims) // create token
	signed, err := token.SignedString([]byte("secret"))
	if err != nil {
		log.Println("Service - SignedToken error : ", err)
	}
	_, errLogin := s.loginRepository.Login(ctx, RequestMapper(user.ID, signed), user.ID)
	if errLogin != nil {
		log.Println("Service - LoginUser error : ", errLogin)
	}

	userId := strconv.FormatUint(uint64(user.ID), 10)

	//Get user By Id
	userData, errGetUser := s.userRepository.GetUserID(ctx, userId)
	if errGetUser != nil {
		log.Println("Service - GetUserId error : ", err)
	}

	return &Response{User: userData}, nil
}
