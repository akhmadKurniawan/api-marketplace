package login

import (
	"app/application/infrastructure"
	"app/middleware"
	"context"
	"errors"
	"log"
	"os"
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
	user, errUser := s.userRepository.GetAllUsername(ctx, req.Username)
	if errUser != nil {
		log.Println("Service - Login error while access username : ", errUser)
		return nil, errUser
	}
	byteDBPass := []byte(user.Password)
	byteReqPass := []byte(req.Password)

	//compare hash password
	if error := bcrypt.CompareHashAndPassword(byteDBPass, byteReqPass); error != nil {
		log.Println("Service - Compare hash error : ", error)
		return nil, error
	}

	//create Claims
	claims := middleware.CreateClaims(uint64(user.ID), user.Username, user.Role, time.Duration(64))
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims) // create token
	signed, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		log.Println("Service - SignedToken error : ", err)
		return nil, err
	}
	_, errLogin := s.loginRepository.Login(ctx, RequestMapper(user.ID, signed), user.ID)
	if errLogin != nil {
		log.Println("Service - LoginUser error : ", errLogin)
		return nil, errLogin
	}

	userId := strconv.FormatUint(uint64(user.ID), 10)

	//Get user By Id
	userData, errGetUser := s.userRepository.GetUserID(ctx, userId)
	if errGetUser != nil {
		log.Println("Service - GetUserId error : ", err)
		return nil, errGetUser
	}
	if errGetUser != nil || userData.Status == "Inactivated" {
		errGetUser = errors.New("please activate your account")
		return nil, errGetUser
	}

	return &Response{User: userData}, nil

}
