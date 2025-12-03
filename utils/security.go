package utils

import ( //looked up imports
	"crypto/sha256"
	"encoding/hex"
	"unicode"
)

// Hash style pssword, takes a password and returns a SHA256 hash.
// Use of a one-way function, so you can't get the original password back.
// Had to look it up, acquired from go.dev
func HashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

// Checks if password is valid, based on out standard
// Standard: Password is at least 20 characters also it has a letter and has a number
func ValidatePassword(password string) bool {
	if len(password) < 20 {
		return false
	}

	hasLetter := false
	hasNumber := false

	for _, ch := range password {
		if unicode.IsLetter(ch) {
			hasLetter = true
		}
		if unicode.IsDigit(ch) {
			hasNumber = true
		}
	}

	return hasLetter && hasNumber
}
