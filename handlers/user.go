package handlers

import (
	"ayobelajar-app-backend/inputs"
	"ayobelajar-app-backend/middleware"
	"ayobelajar-app-backend/responses"
	"ayobelajar-app-backend/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *userHandler {
	return &userHandler{userService}
}

func (uh *userHandler) GetAllUsersHandler(c *gin.Context) {
	users, err := uh.userService.GetUsers()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "200",
			"message": "All users data is empty! Please create new user now!",
		})
		return
	}

	var usersRsps []responses.GetUsersResponse

	for _, user := range users {
		userRsps := responses.ConvertToGetUsersResponse(user)

		usersRsps = append(usersRsps, userRsps)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "All users data is succesfully appeared!",
		"data":    usersRsps,
	})
}

func (uh *userHandler) GetOneUserHandler(c *gin.Context) {
	username := c.Param("username")

	user, err := uh.userService.GetUser(username)

	// fmt.Println(user)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "400",
			"message": "The " + username + " data is not appeared! Please try again!",
		})
		return
	}

	userRsps := responses.ConvertToGetUsersResponse(user)

	c.JSON(http.StatusOK, gin.H{
		"status":  "400",
		"message": "The " + username + " data is succesfully appeared!",
		"data":    userRsps,
	})
}

func (uh *userHandler) CreateNewUserHandler(c *gin.Context) {
	var createUserInput inputs.CreateUserInput

	errCreateUserInput := c.ShouldBindJSON(&createUserInput)

	if errCreateUserInput != nil {
		errorMessages := []string{}

		for _, e := range errCreateUserInput.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": errorMessages,
		})

		return
	}

	_, err := uh.userService.GetUser(createUserInput.UserUsername)

	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "A username data is registered! Please try to find different new username!",
		})
		return
	}

	userUUID, userName, _, _ := middleware.CurrentUser(c)

	current_user_data := map[string]string{
		"current_user_uuid":     userUUID,
		"current_user_username": userName,
	}

	newUser, _ := uh.userService.CreateUser(createUserInput, current_user_data)

	createUserRsps := responses.ConvertToCreateUserResponse(newUser)

	c.JSON(http.StatusOK, gin.H{
		"status":  "400",
		"message": "A new username data is succesfully created!",
		"data":    createUserRsps,
	})
}

func (uh *userHandler) UpdateOneUserHandler(c *gin.Context) {
	username := c.Param("username")

	_, err := uh.userService.GetUser(username)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "400",
			"message": "The " + username + " data is not appeared! Please try again!",
		})
		return
	}

	var updateUserInput inputs.UpdateUserInput

	errUpdateUserInput := c.ShouldBindJSON(&updateUserInput)

	if errUpdateUserInput != nil {
		errorMessages := []string{}

		for _, e := range errUpdateUserInput.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": errorMessages,
		})

		return
	}

	userUUID, userName, _, _ := middleware.CurrentUser(c)

	current_user_data := map[string]string{
		"current_user_uuid":     userUUID,
		"current_user_username": userName,
	}

	updateUser, err := uh.userService.UpdateUser(username, updateUserInput, current_user_data)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "400",
			"message": "The process for updating " + username + " data is failed! Please try again!",
		})
		return
	}

	updateUserRsps := responses.ConvertToUpdateUserResponse(updateUser)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "A username " + username + " data is succesfully updated!",
		"data":    updateUserRsps,
	})
}

func (uh *userHandler) DeleteOneUserHandler(c *gin.Context) {
	username := c.Param("username")

	_, err := uh.userService.GetUser(username)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "400",
			"message": "The " + username + " data is not appeared! Please try again!",
		})
		return
	}

	_, errDeleteUser := uh.userService.DeleteUser(username)

	if errDeleteUser != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "400",
			"message": "The process for deleting " + username + " data is failed! Please try again!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "A username " + username + " data is succesfully deleted!",
	})
}
