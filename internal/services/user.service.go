package services

import (
	"context"

	"github.com/tqlong1609/go_backend_ecommerce/internal/model"
)

type RegisterResponse struct {
	UserID int `json:"userID"`
}

type (
	IUserLogin interface {
		RegisterWithEmail(ctx context.Context, params model.RegisterWithEmailInput) error
		Login(ctx context.Context) error
		Logout(ctx context.Context) error
		VerifyOTP(ctx context.Context, params model.VerifyOTPInput) error
		UpdatePassword(ctx context.Context) error
	}
	IUserInfo interface {
		GetUserInfo(ctx context.Context) error
	}
	IUserAdmin interface {
		GetUserAdmin(ctx context.Context) error
	}
)

var (
	localUserLogin IUserLogin
	localUserInfo  IUserInfo
	localUserAdmin IUserAdmin
)

func InitUserLogin(i IUserLogin) {
	localUserLogin = i
}

func InitUserInfo(i IUserInfo) {
	localUserInfo = i
}

func InitUserAdmin(i IUserAdmin) {
	localUserAdmin = i
}

func UserLogin() IUserLogin {
	if localUserLogin == nil {
		panic("implement not found for interface IUserLogin, forgot register?")
	}
	return localUserLogin
}

func UserInfo() IUserInfo {
	if localUserInfo == nil {
		panic("implement not found for interface IUserInfo, forgot register?")
	}
	return localUserInfo
}

func UserAdmin() IUserAdmin {
	if localUserAdmin == nil {
		panic("implement not found for interface IUserAdmin, forgot register?")
	}
	return localUserAdmin
}
