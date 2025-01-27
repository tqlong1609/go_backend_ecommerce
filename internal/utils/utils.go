package utils

import (
	"math/rand"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

func GenerateSalt() string {
	const saltLength = 16
	const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	salt := make([]byte, saltLength)
	for i := range salt {
		salt[i] = characters[rand.Intn(len(characters))]
	}
	return string(salt)
}

func GeneratePassword(password, salt string) string {
	return password + salt
}

func VerifyPassword(password, hashedPassword, salt string) bool {
	return hashedPassword == GeneratePassword(password, salt)
}

func GenerateToken(userId int32, tokenType string, duration time.Duration) (string, error) {
	// Khởi tạo claims (payload) của token
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(duration).Unix(), // Thời gian hết hạn của token
		"iat":     time.Now().Unix(),               // Thời gian phát hành token
		"type":    tokenType,                       // Loại token (access/refresh)
	}

	// Tạo token với phương thức ký (HS256)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Ký token bằng secret key
	secretKey := []byte(constants.JWT_SECRET_KEY) // Thay bằng khóa bí mật thực tế của bạn
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err // Trả về lỗi nếu có
	}

	return tokenString, nil
}
