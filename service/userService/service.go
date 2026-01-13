package userservice

import (
	"fmt"
	"room-reserve/entity"
	"room-reserve/pkg/phonenumber"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	RegisterUser(u entity.User) (entity.User, error)
}

type Service struct {
	repo Repository
}

type RegisterRequest struct {
	Name        string
	PhoneNumher string
}

type RegisterResponse struct {
	entity.User
}

func (s Service) Register(req RegisterRequest) (RegisterResponse, error) {
	// TODO - we should verify phone number by verification code

	// validate phone number
	if !phonenumber.IsValid(req.PhoneNumher) {
		return RegisterResponse{}, nil
	}

	// validate unique phone number
	if isUnique, err := s.repo.IsPhoneNumberUnique(req.PhoneNumher); err != nil || !isUnique {
		if err != nil {
			return RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)
		}

		if !isUnique {
			return RegisterResponse{}, fmt.Errorf("phone number is not unique")
		}
	}

	// validate name
	if len(req.Name) < 3 {
		return RegisterResponse{}, fmt.Errorf("name should be greater then 3 character")
	}

	// create new user
	user := entity.User{
		ID:          0,
		PhoneNumber: req.Name,
		Name:        req.Name,
	}
	createdUser, err := s.repo.RegisterUser(user)
	if err != nil {
		return RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	// return user
	return RegisterResponse{
		User: createdUser,
	}, nil
}
