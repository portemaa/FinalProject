package utils

import "math"

// Checks how strong a password is based on length and variety of chars
func CheckPasswordStrength(password string) float64 {
	charsetSize := CheckCharType(password)
	length := len(password)

	if charsetSize == 0 || length == 0 {
		return 0
	}

	// Formula to check strength
	return math.Log2(float64(charsetSize)) * float64(length)
}

// Checks how many types of chars are used
func CheckCharType(password string) int {
	size := 0
	hasLower, hasUpper, hasDigit, hasSymbol := false, false, false, false

	for _, ch := range password {
		switch {
		case 'a' <= ch && ch <= 'z':
			hasLower = true
		case 'A' <= ch && ch <= 'Z':
			hasUpper = true
		case '0' <= ch && ch <= '9':
			hasDigit = true
		default:
			hasSymbol = true
		}
	}

	if hasLower {
		size += 26
	}
	if hasUpper {
		size += 26
	}
	if hasDigit {
		size += 10
	}
	if hasSymbol {
		size += 32
	}

	return size
}
