package hash

import "golang.org/x/crypto/bcrypt"

func Generate(password *string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(*password), 10)
	return string(bytes), err
}

func CheckHash(password, hash *string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(*hash), []byte(*password))
	return err == nil
}
