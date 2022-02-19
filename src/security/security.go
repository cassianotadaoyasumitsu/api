package security

import "golang.org/x/crypto/bcrypt"

// Create a hash for the password
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// Verify hash / password
func Verify(passHash, passString string) error {
	return bcrypt.CompareHashAndPassword([]byte(passHash), []byte(passString))
}
