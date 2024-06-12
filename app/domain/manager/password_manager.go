package manager

import "golang.org/x/crypto/bcrypt"

type PasswordManager interface {
	CreateHashPassword(password string) (string, error)
	CheckHashPassword(password string, hash string) bool
}

type passwordManagerImplementation struct {
	bcryptCost int
}

func (m *passwordManagerImplementation) CreateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (m *passwordManagerImplementation) CheckHashPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//func HashPassword(password string) (string, error) {
//	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
//	return string(bytes), err
//}
//
//func CheckPasswordHash(password, hash string) bool {
//	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
//	return err == nil
//}

func NewPasswordManager(bcryptCost int) PasswordManager {
	return &passwordManagerImplementation{bcryptCost: bcryptCost}
}