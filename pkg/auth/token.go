package auth

func ValidateToken(token, expectedToken string) bool {
	return token == expectedToken
}
