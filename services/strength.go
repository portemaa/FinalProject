package services

import "password-checker/utils"

// Looks at the password and returns a title,
// either "Weak", "Moderate", or "Strong"
func GetPasswordStrength(password string) string {
	score := utils.CheckPasswordStrength(password)

	// Decide score based on password strength
	if score < 28 {
		return "Weak"
	} else if score < 36 {
		return "Moderate"
	} else {
		return "Strong"
	}
}
