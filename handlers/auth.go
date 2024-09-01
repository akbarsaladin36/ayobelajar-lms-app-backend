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

type authHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *authHandler {
	return &authHandler{authService}
}

func (ah *authHandler) RegisterUserHandler(c *gin.Context) {
	var registerInput inputs.RegisterInput

	errRegisterInput := c.ShouldBindJSON(&registerInput)

	if errRegisterInput != nil {
		errorMessages := []string{}

		for _, e := range errRegisterInput.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": errorMessages,
		})

		return
	}

	_, error := ah.authService.GetUsername(registerInput.UserUsername)

	if error == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "The username is exist!, Please try to find new different username!",
		})
		return
	}

	newUser, err := ah.authService.RegisterUser(registerInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Process registering user is failed! Please try again!",
		})
		return
	}

	registerRsps := responses.ConvertToAuthRegisterUser(newUser)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "A new user is succesfully registered!",
		"data":    registerRsps,
	})
}

func (ah *authHandler) LoginUserHandler(c *gin.Context) {
	var loginInput inputs.LoginInput

	errLoginInput := c.ShouldBindJSON(&loginInput)

	if errLoginInput != nil {
		errorMessages := []string{}

		for _, e := range errLoginInput.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": errorMessages,
		})

		return
	}

	_, error := ah.authService.GetUsername(loginInput.UserUsername)

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "The username is not exist!, Please try to find new different username!",
		})
		return
	}

	loginUser, errorLoginUser := ah.authService.LoginUser(loginInput)

	if errorLoginUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "The password is not matched! Please try again!",
		})
		return
	}

	tokenString, errorTokenString := middleware.GenerateJWTAuthentication(loginUser.UserUUID, loginUser.UserUsername, loginUser.UserEmail)

	if errorTokenString != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "The token is not succesfully generated! Please try again!",
		})
		return
	}

	loginResponse := responses.ConvertToAuthLoginUser(loginUser, tokenString)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "The user is succesfully login!",
		"data":    loginResponse,
	})

}
