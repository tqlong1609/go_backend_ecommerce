package implements

import (
	"context"
	"fmt"
	"time"

	"github.com/tqlong1609/go_backend_ecommerce/global"
	"github.com/tqlong1609/go_backend_ecommerce/internal/database"
	"github.com/tqlong1609/go_backend_ecommerce/internal/model"
	"github.com/tqlong1609/go_backend_ecommerce/internal/utils"
	"go.uber.org/zap/zapcore"
)

type UserLoginImplement struct {
	dbQueries *database.Queries
}

func InitUserLoginImplement(dbQueries *database.Queries) *UserLoginImplement {
	return &UserLoginImplement{
		dbQueries: dbQueries,
	}
}

func (ul *UserLoginImplement) RegisterWithEmail(ctx context.Context, params model.RegisterWithEmailInput) error {
	if params.Email == "" {
		return fmt.Errorf("email is empty")
	}

	// check email valid
	if !utils.IsValidEmail(params.Email) {
		return fmt.Errorf("email is invalid")
	}

	// check email exist
	_, err := ul.dbQueries.FindUserByEmail(ctx, params.Email)
	if err == nil {
		return fmt.Errorf("email is exist")
	}

	// Lưu OTP vào cơ sở dữ liệu
	otp := utils.GenerateOTP()
	err = ul.dbQueries.InsertOrUpdateOtp(ctx, database.InsertOrUpdateOtpParams{
		VerifyKey: params.Email,
		VerifyOtp: otp,
	})

	if err != nil {
		return err
	}

	global.Logger.Info("Sent OTP successfully", zapcore.Field{
		Key:    "otp",
		Type:   zapcore.StringType,
		String: otp,
	}, zapcore.Field{Key: "email", Type: zapcore.StringType, String: params.Email})

	return nil
}

func (ul *UserLoginImplement) Login(ctx context.Context) error {
	return nil
}

func (ul *UserLoginImplement) Logout(ctx context.Context) error {
	return nil
}

func (ul *UserLoginImplement) VerifyOTP(ctx context.Context, params model.VerifyOTPInput) error {
	if params.Email == "" {
		return fmt.Errorf("email is empty")
	}

	// check email valid
	if !utils.IsValidEmail(params.Email) {
		return fmt.Errorf("email is invalid")
	}

	otpRecord, err := ul.dbQueries.GetOtpRecord(ctx, params.Email)
	if err != nil {
		return err
	}
	if otpRecord.VerifyOtp == "" {
		return fmt.Errorf("otp is empty")
	}

	// Kiểm tra xem OTP có bị khóa không
	if otpRecord.LockUntil.Valid && otpRecord.LockUntil.Time.After(time.Now()) {
		return fmt.Errorf("otp is locked")
	}

	// Kiểm tra mã OTP không hợp lệ
	if otpRecord.VerifyOtp != params.OTP {
		failedAttempts := otpRecord.FailedAttempts + 1

		// Cập nhật số lần nhập sai
		ul.dbQueries.UpdateFailedAttempts(ctx, database.UpdateFailedAttemptsParams{
			VerifyKey:      params.Email,
			FailedAttempts: failedAttempts,
		})

		// Khóa OTP nếu vượt giới hạn >= 5, 10 phút
		if failedAttempts >= 5 {
			lockUntil := otpRecord.LockUntil
			lockUntil.Time = time.Now().Add(time.Minute * 10)
			ul.dbQueries.LockOtp(ctx, database.LockOtpParams{
				VerifyKey: params.Email,
				LockUntil: lockUntil,
			})
			return fmt.Errorf("too many failed attempts. otp locked for 10 minutes")
		}
		return fmt.Errorf("otp is invalid")
	}

	// OTP hợp lệ, cập nhật trạng thái và reset số lần nhập sai
	ul.dbQueries.VerifySuccessOtp(ctx, params.Email)
	return nil
}

// complete registration
func (ul *UserLoginImplement) CompleteRegistration(ctx context.Context, params model.CompleteRegistrationInput) error {
	if params.Email == "" || params.Password == "" {
		return fmt.Errorf("email or password is empty")
	}

	// Kiểm tra xem tài khoản đã tồn tại chưa
	_, err := ul.dbQueries.FindUserByEmail(ctx, params.Email)
	if err == nil {
		return fmt.Errorf("email is exist")
	}

	// Tạo salt và hash mật khẩu
	salt := utils.GenerateSalt()
	passwordHash := utils.GeneratePassword(params.Password, salt)

	// Thêm người dùng vào cơ sở dữ liệu
	userInfoBase, err := ul.dbQueries.AddUserInfoBase(ctx, database.AddUserInfoBaseParams{
		UserAccount:  params.Email,
		UserPassword: passwordHash,
		UserSalt:     salt,
	})
	if err != nil {
		return err
	}

	_, err = ul.dbQueries.InsertUserInfo(ctx, database.InsertUserInfoParams{
		UserID:      userInfoBase.UserID,
		UserAccount: params.Email,
		UserName:    utils.ToNullString(params.Name),
	})
	if err != nil {
		return err
	}
	return nil
}

func (ul *UserLoginImplement) UpdatePassword(ctx context.Context) error {
	return nil
}
