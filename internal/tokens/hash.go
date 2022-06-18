package tokens

import "golang.org/x/crypto/bcrypt"

func HashToken(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hash), err
}

func CheckTokenHash(token, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(token))
	return err == nil
}
