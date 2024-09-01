package inputs

type RegisterInput struct {
	UserUsername string `json:"user_username"`
	UserEmail    string `json:"user_email"`
	UserPassword string `json:"user_password"`
}

type LoginInput struct {
	UserUsername string `json:"user_username"`
	UserPassword string `json:"user_password"`
}
