package service

import (
	model "cat-social/models"
	"cat-social/models/dto/request"
	"cat-social/models/dto/response"
	repository "cat-social/repositories"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Create(signUpRequest request.SignupRequest) (response.SignUpResponse, error)
	Login(loginRequest request.SignInRequest) (string, error)
	GenerateToken(user model.User) (string, error)
}
type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) Create(signUpRequest request.SignupRequest) (response.SignUpResponse, error) {
	//hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(signUpRequest.Password), 10)

	if err != nil {
		return response.SignUpResponse{}, err
	}

	var isEmailExist bool = s.repository.EmailIsExist(signUpRequest.Email)
	fmt.Println(isEmailExist)
	if isEmailExist {
		fmt.Println("hitted error email redudant")
		return response.SignUpResponse{}, errors.New("EMAIL ALREADY EXIST")
	}

	//save user
	user := model.User{
		Email:    signUpRequest.Email,
		Name:     signUpRequest.Name,
		Password: string(hash),
	}

	newUser, err := s.repository.Create(user)
	if err != nil {
		return response.SignUpResponse{}, err
	}

	//sign token
	token, err := s.GenerateToken(newUser)
	if err != nil {
		return response.SignUpResponse{}, err
	}

	// Create and return SignUpResponse
	signUpResponse := response.SignUpResponse{
		Name:        newUser.Name,
		Email:       newUser.Email,
		AccessToken: token,
	}

	return signUpResponse, nil
}

func (s *userService) Login(loginRequest request.SignInRequest) (string, error) {
	//get user
	user, err := s.repository.FindByEmail(loginRequest.Email)

	if err != nil {
		return "", err
	} else if user.ID == 0 {
		return "", errors.New("Invalid email or password")
	}

	//compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))

	if err != nil {
		return "", err
	}

	//sign token
	token, err := s.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) GenerateToken(user model.User) (string, error) {
	//sign token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
