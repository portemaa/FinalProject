package services

// Function gives you simple suggestions to make your password stronger
func ImprovePassword(password string) string {
	strength := GetPasswordStrength(password)

	if strength == "Weak" {
		return "Try making your password longer and include numbers, uppercase letters, and symbols."
	} else if strength == "Moderate" {
		return "Your password is fine, but adding more complex characters(symbols or uppercase letters) will help."
	} else {
		return "Perfect, Your password is strong."
	}
}
