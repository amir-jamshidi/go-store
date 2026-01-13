package phonenumber

import "regexp"

func IsValid(phoneNumber string) bool {
	// TODO - This fn not support +98
	return regexp.MustCompile(`^09[\d]{9}$`).MatchString(phoneNumber)
}
