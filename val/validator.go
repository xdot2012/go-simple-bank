package val

import (
	"fmt"
	"net/mail"
	"regexp"
)

const (
	MIN_USERNAME_LENGTH = 3
	MAX_USERNAME_LENGTH = 100
	MIN_FULLNAME_LENGTH = 3
	MAX_FULLNAME_LENGTH = 100
	MIN_PASSWORD_LENGTH = 6
	MAX_PASSWORD_LENGTH = 100
	MIN_EMAIL_LENGTH    = 3
	MAX_EMAIL_LENGTH    = 200
)

var (
	isValidUsername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
	isValidFullName = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
)

func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("must contain from %d-%d characters", minLength, maxLength)
	}
	return nil
}

func ValidateUsername(value string) error {
	if err := ValidateString(value, MIN_USERNAME_LENGTH, MAX_USERNAME_LENGTH); err != nil {
		return err
	}

	if !isValidUsername(value) {
		return fmt.Errorf("must contain only lowercase letters, digits or underscore")
	}
	return nil
}

func ValidateFullName(value string) error {
	if err := ValidateString(value, MIN_FULLNAME_LENGTH, MAX_FULLNAME_LENGTH); err != nil {
		return err
	}

	if !isValidFullName(value) {
		return fmt.Errorf("must contain only letters, or spaces")
	}
	return nil
}

func ValidatePassword(value string) error {
	return ValidateString(value, MIN_PASSWORD_LENGTH, MAX_PASSWORD_LENGTH)
}

func ValidateEmail(value string) error {
	if err := ValidateString(value, MIN_EMAIL_LENGTH, MAX_EMAIL_LENGTH); err != nil {
		return err
	}
	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("email invalid")
	}
	return nil
}
