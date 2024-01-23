package handlers

import (
	"net/mail"
	"regexp"
)

func isEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func isValid(input string) bool {
	r := regexp.MustCompile("^[a-zA-Z0-9]{3,}$")
	return r.MatchString(input)
}
