package model

type RegisterInput struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type RegisterOutput struct {
	UserID int32 `json:"userID"`
}

type RegisterWithEmailInput struct {
	Email string `json:"email"`
}

type VerifyOTPInput struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}

type CompleteRegistrationInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginOutput struct {
	UserID       int32  `json:"userID"`
	AccessToken  string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}
