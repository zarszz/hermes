package entity

type UserRegistrationInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type UserLoginInput struct {
	Email string `json:"email"`
	Password string `json:"password"`
}
