package security

import (
	"regexp"
	"strings"
	"unicode"
)

// PasswordValidationResult represents the result of password validation
type PasswordValidationResult struct {
	IsValid bool     `json:"is_valid"`
	Errors  []string `json:"errors"`
	Score   int      `json:"score"` // 0-4 (weak to very strong)
}

// ValidatePassword validates password strength and returns detailed results
func ValidatePassword(password string) PasswordValidationResult {
	result := PasswordValidationResult{
		IsValid: true,
		Errors:  []string{},
		Score:   0,
	}

	// Check minimum length
	if len(password) < 8 {
		result.IsValid = false
		result.Errors = append(result.Errors, "Password must be at least 8 characters long")
	} else {
		result.Score++
	}

	// Check for uppercase letters
	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		result.IsValid = false
		result.Errors = append(result.Errors, "Password must contain at least one uppercase letter")
	} else {
		result.Score++
	}

	// Check for lowercase letters
	if !regexp.MustCompile(`[a-z]`).MatchString(password) {
		result.IsValid = false
		result.Errors = append(result.Errors, "Password must contain at least one lowercase letter")
	} else {
		result.Score++
	}

	// Check for numbers
	if !regexp.MustCompile(`[0-9]`).MatchString(password) {
		result.IsValid = false
		result.Errors = append(result.Errors, "Password must contain at least one number")
	} else {
		result.Score++
	}

	// Check for special characters
	if !regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`).MatchString(password) {
		result.Errors = append(result.Errors, "Password should contain at least one special character for better security")
	} else {
		result.Score++
	}

	// Check for common weak patterns
	weakPatterns := []string{"123456", "password", "qwerty", "abc123", "admin"}
	lowerPassword := strings.ToLower(password)
	for _, pattern := range weakPatterns {
		if strings.Contains(lowerPassword, pattern) {
			result.IsValid = false
			result.Errors = append(result.Errors, "Password contains common weak patterns")
			result.Score = 0
			break
		}
	}

	// Check for sequential characters
	if hasSequentialChars(password) {
		result.Errors = append(result.Errors, "Password contains sequential characters (e.g., 123, abc)")
		if result.Score > 0 {
			result.Score--
		}
	}

	// Check for repeated characters
	if hasRepeatedChars(password, 3) {
		result.Errors = append(result.Errors, "Password contains too many repeated characters")
		if result.Score > 0 {
			result.Score--
		}
	}

	return result
}

// ValidateEmail validates email format and checks for common security issues
func ValidateEmail(email string) (bool, []string) {
	errors := []string{}

	// Basic format validation
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		errors = append(errors, "Invalid email format")
		return false, errors
	}

	// Check for suspicious patterns
	if strings.Contains(email, "..") {
		errors = append(errors, "Email contains suspicious patterns")
	}

	// Check length
	if len(email) > 254 {
		errors = append(errors, "Email too long")
		return false, errors
	}

	// Check local part length
	parts := strings.Split(email, "@")
	if len(parts[0]) > 64 {
		errors = append(errors, "Email local part too long")
		return false, errors
	}

	return len(errors) == 0, errors
}

// SanitizeInput removes potentially dangerous characters from user input
func SanitizeInput(input string) string {
	// Remove null bytes
	input = strings.ReplaceAll(input, "\x00", "")

	// Remove control characters except tab, newline, and carriage return
	var result strings.Builder
	for _, r := range input {
		if unicode.IsControl(r) && r != '\t' && r != '\n' && r != '\r' {
			continue
		}
		result.WriteRune(r)
	}

	return strings.TrimSpace(result.String())
}

// ValidateJWTSecret validates JWT secret strength
func ValidateJWTSecret(secret string) (bool, []string) {
	errors := []string{}

	if len(secret) < 32 {
		errors = append(errors, "JWT secret must be at least 32 characters long")
	}

	if len(secret) < 64 {
		errors = append(errors, "JWT secret should be at least 64 characters for optimal security")
	}

	// Check for common weak secrets
	weakSecrets := []string{"secret", "jwt_secret", "your-secret-key", "mysecret", "jwt"}
	lowerSecret := strings.ToLower(secret)
	for _, weak := range weakSecrets {
		if strings.Contains(lowerSecret, weak) {
			errors = append(errors, "JWT secret contains common weak patterns")
			break
		}
	}

	// Check for sufficient entropy (basic check)
	if !hasGoodEntropy(secret) {
		errors = append(errors, "JWT secret appears to have low entropy")
	}

	return len(errors) == 0, errors
}

// hasSequentialChars checks for sequential character patterns
func hasSequentialChars(s string) bool {
	s = strings.ToLower(s)
	for i := 0; i < len(s)-2; i++ {
		if s[i]+1 == s[i+1] && s[i+1]+1 == s[i+2] {
			return true
		}
		if s[i]-1 == s[i+1] && s[i+1]-1 == s[i+2] {
			return true
		}
	}
	return false
}

// hasRepeatedChars checks for repeated character patterns
func hasRepeatedChars(s string, maxRepeats int) bool {
	charCount := make(map[rune]int)
	for _, char := range s {
		charCount[char]++
		if charCount[char] >= maxRepeats {
			return true
		}
	}
	return false
}

// hasGoodEntropy performs a basic entropy check
func hasGoodEntropy(s string) bool {
	if len(s) < 16 {
		return false
	}

	charTypes := 0
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(s)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(s)
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(s)
	hasSpecial := regexp.MustCompile(`[^a-zA-Z0-9]`).MatchString(s)

	if hasLower {
		charTypes++
	}
	if hasUpper {
		charTypes++
	}
	if hasDigit {
		charTypes++
	}
	if hasSpecial {
		charTypes++
	}

	return charTypes >= 3
}
