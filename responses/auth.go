package responses

import (
	"ayobelajar-app-backend/models"
)

type RegisterUserResponse struct {
	UserUsername string `json:"user_username"`
	UserEmail    string `json:"user_email"`
	UserRole     string `json:"user_role"`
	UserStatusCd string `json:"user_status_cd"`
}

type LoginUserResponse struct {
	UserUsername string `json:"user_username"`
	UserEmail    string `json:"user_email"`
	UserToken    string `json:"user_token"`
}

func ConvertToAuthRegisterUser(authUserRsps models.User) RegisterUserResponse {
	return RegisterUserResponse{
		UserUsername: authUserRsps.UserUsername,
		UserEmail:    authUserRsps.UserEmail,
		UserRole:     authUserRsps.UserRole,
		UserStatusCd: authUserRsps.UserStatusCd,
	}
}

func ConvertToAuthLoginUser(authUserRsps models.User, tokenString string) LoginUserResponse {
	return LoginUserResponse{
		UserUsername: authUserRsps.UserUsername,
		UserEmail:    authUserRsps.UserEmail,
		UserToken:    tokenString,
	}
}
