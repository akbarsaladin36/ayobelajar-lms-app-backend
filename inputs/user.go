package inputs

type CreateUserInput struct {
	UserUsername    string `json:"user_username"`
	UserEmail       string `json:"user_email"`
	UserPassword    string `json:"user_password"`
	UserFirstName   string `json:"user_first_name"`
	UserLastName    string `json:"user_last_name"`
	UserAddress     string `json:"user_address"`
	UserPhoneNumber string `json:"user_phone_number"`
}

type UpdateUserInput struct {
	UserFirstName   string `json:"user_first_name"`
	UserLastName    string `json:"user_last_name"`
	UserAddress     string `json:"user_address"`
	UserPhoneNumber string `json:"user_phone_number"`
	UserRole        string `json:"user_role"`
}
