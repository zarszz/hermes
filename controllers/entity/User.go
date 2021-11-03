package entity

type UserRegistrationInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}
type UserUpdateInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserLoginInput struct {
	Email string `json:"email"`
	Password string `json:"password"`
}
