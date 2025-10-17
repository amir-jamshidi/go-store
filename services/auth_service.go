package services

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"room-reserve/models"
	"room-reserve/repositories"
	"room-reserve/utils"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type AuthService struct {
	userRepo *repositories.UserRepository
	otpRepo  *repositories.OTPRepository
}

func NewAuthService(userRepo *repositories.UserRepository, otpRepo *repositories.OTPRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		otpRepo:  otpRepo,
	}
}

func (s *AuthService) SendOTP(ctx context.Context, phone string) (string, error) {
	if !s.isValidPhoneNumber(phone) {
		return "", errors.New("شماره تلفن نامعتبر است")
	}

	user, err := s.userRepo.FindByPhone(ctx, phone)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			user = &models.User{
				Phone:     phone,
				Role:      "user",
				Score:     0,
				CreateAt:  time.Now(),
				UpdatedAt: time.Now(),
			}

			if err := s.userRepo.Create(ctx, user); err != nil {
				return "", fmt.Errorf("خطا در ایجاد کاربر: %w", err)
			}
		} else {
			return "", fmt.Errorf("خطا در جستجوی کاربر: %w", err)
		}
	}

	s.otpRepo.DeleteByUserID(ctx, user.ID)

	otpCode := s.generateOTP()

	newOTP := &models.OTP{
		Code:      otpCode,
		UserID:    user.ID,
		TryCount:  0,
		ExpireAt:  time.Now().Add(5 * time.Minute),
		CreatedAt: time.Now(),
	}

	if err := s.otpRepo.Create(ctx, newOTP); err != nil {
		return "", fmt.Errorf("خطا در ذخیره OTP: %w", err)
	}

	// TODO: SEND SMS
	fmt.Printf("OTP for %s: %s\n", phone, otpCode)

	return otpCode, nil
}

func (s *AuthService) VerifyOTP(ctx context.Context, phone string, code string) (*models.User, string, error) {
	user, err := s.userRepo.FindByPhone(ctx, phone)
	if err != nil {
		return nil, "", errors.New("کاربر یافت نشد")
	}

	otp, err := s.otpRepo.FindByUserID(ctx, user.ID)
	if err != nil {
		return nil, "", errors.New("کد OTP یافت نشد یا منقضی شده")
	}

	if time.Now().After(otp.ExpireAt) {
		s.otpRepo.Delete(ctx, otp.ID)
		return nil, "", errors.New("کد OTP منقضی شده است")
	}

	if otp.TryCount >= 3 {
		s.otpRepo.Delete(ctx, otp.ID)
		return nil, "", errors.New("تعداد تلاش بیش از حد. لطفاً کد جدید درخواست کنید")
	}

	if otp.Code != code {
		s.otpRepo.IncrementTryCount(ctx, otp.ID)
		return nil, "", errors.New("کد OTP نامعتبر است")
	}

	s.otpRepo.Delete(ctx, otp.ID)

	s.userRepo.UpdateLastLogin(ctx, user.ID)

	token, err := utils.GenerateJWT(user.ID.Hex(), user.Role)
	if err != nil {
		return nil, "", fmt.Errorf("خطا در تولید توکن: %w", err)
	}

	return user, token, nil

}

func (s *AuthService) isValidPhoneNumber(phone string) bool {
	matched, _ := regexp.MatchString(`^09[0-9]{9}$`, phone)
	return matched
}

func (s *AuthService) generateOTP() string {
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}
