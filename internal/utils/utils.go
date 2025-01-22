package utils

import (
	"math/rand"
	"regexp"

	"github.com/tqlong1609/go_backend_ecommerce/internal/constants"
)

func IsValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}`)
	return emailRegex.MatchString(email)
}

func GenerateOTP() string {
	const digits = "0123456789"
	length := constants.OTP_LENGTH
	otp := make([]byte, length)

	for i := 0; i < length; i++ {
		otp[i] = digits[rand.Intn(len(digits))]
	}

	return string(otp)
}
