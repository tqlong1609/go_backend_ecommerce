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
