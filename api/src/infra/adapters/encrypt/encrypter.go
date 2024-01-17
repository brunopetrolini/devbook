package encrypter

import "golang.org/x/crypto/bcrypt"

// Hash is a function to encrypt an input string
func Hash(str string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
}

// Compare is a function to verify an input string with a hash
func Compare(hash string, str string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))
}
