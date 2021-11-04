package utils

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(password string) (*string, error)  {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	newPassword := string(hashedPassword)
	return &newPassword, nil
}

func ComparePassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
