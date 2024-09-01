package responses

import "ayobelajar-app-backend/models"

type GetUsersResponse struct {
	UserUsername    string `json:"user_username"`
	UserEmail       string `json:"user_email"`
	UserFirstName   string `json:"user_first_name"`
	UserLastName    string `json:"user_last_name"`
	UserAddress     string `json:"user_address"`
	UserPhoneNumber string `json:"user_phone_number"`
}

type CreateUserResponse struct {
	UserUsername    string `json:"user_username"`
	UserEmail       string `json:"user_email"`
	UserPassword    string `json:"user_password"`
	UserFirstName   string `json:"user_first_name"`
	UserLastName    string `json:"user_last_name"`
	UserAddress     string `json:"user_address"`
	UserPhoneNumber string `json:"user_phone_number"`
}

type UpdateUserResponse struct {
	UserFirstName   string `json:"user_first_name"`
	UserLastName    string `json:"user_last_name"`
	UserAddress     string `json:"user_address"`
	UserPhoneNumber string `json:"user_phone_number"`
	UserRole        string `json:"user_role"`
}

func ConvertToGetUsersResponse(allUsersRsps models.User) GetUsersResponse {
	return GetUsersResponse{
		UserUsername:    allUsersRsps.UserUsername,
		UserEmail:       allUsersRsps.UserEmail,
		UserFirstName:   allUsersRsps.UserFirstName,
		UserLastName:    allUsersRsps.UserLastName,
		UserAddress:     allUsersRsps.UserAddress,
		UserPhoneNumber: allUsersRsps.UserPhoneNumber,
	}
}

func ConvertToCreateUserResponse(createUserRsps models.User) CreateUserResponse {
	return CreateUserResponse{
		UserUsername:    createUserRsps.UserUsername,
		UserEmail:       createUserRsps.UserEmail,
		UserPassword:    createUserRsps.UserPassword,
		UserFirstName:   createUserRsps.UserFirstName,
		UserLastName:    createUserRsps.UserLastName,
		UserAddress:     createUserRsps.UserAddress,
		UserPhoneNumber: createUserRsps.UserPhoneNumber,
	}
}

func ConvertToUpdateUserResponse(updateUserRsps models.User) UpdateUserResponse {
	return UpdateUserResponse{
		UserFirstName:   updateUserRsps.UserFirstName,
		UserLastName:    updateUserRsps.UserLastName,
		UserAddress:     updateUserRsps.UserAddress,
		UserPhoneNumber: updateUserRsps.UserPhoneNumber,
		UserRole:        updateUserRsps.UserRole,
	}
}
