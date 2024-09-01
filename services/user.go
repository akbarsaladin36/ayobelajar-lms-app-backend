package services

import (
	"ayobelajar-app-backend/helpers"
	"ayobelajar-app-backend/inputs"
	"ayobelajar-app-backend/models"
	"ayobelajar-app-backend/repositories"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetUsers() ([]models.User, error)
	GetUser(username string) (models.User, error)
	CreateUser(userInput inputs.CreateUserInput, currentUser map[string]string) (models.User, error)
	UpdateUser(username string, updateUserInput inputs.UpdateUserInput, currentUser map[string]string) (models.User, error)
	DeleteUser(username string) (models.User, error)
}

type userService struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) *userService {
	return &userService{repository}
}

func (us *userService) GetUsers() ([]models.User, error) {
	users, err := us.repository.GetAll()

	return users, err
}

func (us *userService) GetUser(username string) (models.User, error) {
	user, err := us.repository.GetUser(username)

	return user, err
}

func (us *userService) CreateUser(createUserInput inputs.CreateUserInput, currentUser map[string]string) (models.User, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(createUserInput.UserPassword), 14)

	convertUsernameToUUID := strings.ReplaceAll(helpers.ConvertToUUID(createUserInput.UserUsername), "-", "")

	admin_uuid := currentUser["current_user_uuid"]
	admin_username := currentUser["current_user_username"]

	user := models.User{
		UserUUID:            convertUsernameToUUID,
		UserUsername:        createUserInput.UserUsername,
		UserEmail:           createUserInput.UserEmail,
		UserPassword:        string(hashedPassword),
		UserRole:            "user",
		UserStatusCd:        "active",
		UserCreatedDate:     time.Now(),
		UserCreatedUserUuid: admin_uuid,
		UserCreatedUsername: admin_username,
	}

	newUser, err := us.repository.Create(user)

	return newUser, err
}

func (us *userService) UpdateUser(username string, updateUserInput inputs.UpdateUserInput, currentUser map[string]string) (models.User, error) {
	checkUser, _ := us.repository.GetUser(username)

	admin_uuid := currentUser["current_user_uuid"]
	admin_username := currentUser["current_user_username"]

	checkUser.UserFirstName = updateUserInput.UserFirstName
	checkUser.UserLastName = updateUserInput.UserLastName
	checkUser.UserAddress = updateUserInput.UserAddress
	checkUser.UserPhoneNumber = updateUserInput.UserPhoneNumber
	checkUser.UserRole = updateUserInput.UserRole
	checkUser.UserUpdatedDate = time.Now()
	checkUser.UserUpdatedUserUuid = admin_uuid
	checkUser.UserUpdatedUsername = admin_username

	updateUser, err := us.repository.Update(checkUser)

	return updateUser, err
}

func (us *userService) DeleteUser(username string) (models.User, error) {
	checkUser, _ := us.repository.GetUser(username)

	deleteUser, err := us.repository.Delete(checkUser)

	return deleteUser, err
}
