package services

import (
	"ayobelajar-app-backend/helpers"
	"ayobelajar-app-backend/inputs"
	"ayobelajar-app-backend/models"
	"ayobelajar-app-backend/repositories"
	"fmt"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	GetUsername(username string) (models.User, error)
	RegisterUser(registerInput inputs.RegisterInput) (models.User, error)
	LoginUser(loginInput inputs.LoginInput) (models.User, error)
}

type authService struct {
	authRepository repositories.AuthRepository
}

func NewAuthService(authRepository repositories.AuthRepository) *authService {
	return &authService{authRepository}
}

func (as *authService) GetUsername(username string) (models.User, error) {
	checkUser, err := as.authRepository.GetUserUsername(username)

	return checkUser, err
}

func (as *authService) RegisterUser(registerInput inputs.RegisterInput) (models.User, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(registerInput.UserPassword), 14)

	convertUsernameToUUID := strings.ReplaceAll(helpers.ConvertToUUID(registerInput.UserUsername), "-", "")

	user := models.User{
		UserUUID:            convertUsernameToUUID,
		UserUsername:        registerInput.UserUsername,
		UserEmail:           registerInput.UserEmail,
		UserPassword:        string(hashedPassword),
		UserRole:            "user",
		UserStatusCd:        "active",
		UserCreatedDate:     time.Now(),
		UserCreatedUserUuid: convertUsernameToUUID,
		UserCreatedUsername: registerInput.UserUsername,
	}

	newUser, err := as.authRepository.CreateUser(user)

	return newUser, err
}

func (as *authService) LoginUser(loginInput inputs.LoginInput) (models.User, error) {
	checkUser, err := as.authRepository.GetUserUsername(loginInput.UserUsername)

	if err != nil {
		fmt.Println("The username is not exist! Please try again!")
	}

	checkPassword := bcrypt.CompareHashAndPassword([]byte(checkUser.UserPassword), []byte(loginInput.UserPassword))

	if checkPassword != nil {
		fmt.Println("The password does not match! Please try again!")
	}

	return checkUser, err
}
