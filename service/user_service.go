package service

import (
	"chapter3-sesi2/dto"
	"chapter3-sesi2/entity"
	"chapter3-sesi2/pkg/errs"
	"chapter3-sesi2/pkg/helpers"
	"chapter3-sesi2/repository/user_repository"
	"fmt"
	"net/http"
)

type UserService interface {
	CreateNewUser(payload dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr)
	Login(payload dto.NewUserRequest) (*dto.LoginResponse, errs.MessageErr)
}
type userService struct {
	userRepo user_repository.UserRepository
}

// CreateNewUser implements UserService
func (u *userService) CreateNewUser(payload dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)
	if err != nil {
		return nil, err
	}
	userEntity := entity.User{
		Email:    payload.Email,
		Password: payload.Password,
	}
	err = userEntity.HashPassword()
	if err != nil {
		return nil, err
	}
	err = u.userRepo.CreateNewUser(userEntity)
	if err != nil {
		return nil, err
	}
	response := dto.NewUserResponse{
		Result:     "success",
		Message:    "user registered successfully",
		StatusCode: http.StatusCreated,
	}
	return &response, nil
}

// Login implements UserService
func (u *userService) Login(payload dto.NewUserRequest) (*dto.LoginResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)
	if err != nil {
		return nil, err
	}
	user, err := u.userRepo.GetUserByEmail(payload.Email)
	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, errs.NewUnauthenticatedError("invalid email/password")
		}
		return nil, err
	}
	isValidPassword := user.ComparePassword(payload.Password)
	if !isValidPassword {
		return nil, errs.NewUnauthenticatedError("invalid email/password")
	}
	fmt.Println("user =>", user)
	response := dto.LoginResponse{
		Result:     "success",
		Message:    "logged in successfully",
		StatusCode: http.StatusOK,
		Data: dto.TokenResponse{
			Token: user.GenerateToken(),
		},
	}
	return &response, nil
}

func NewUserService(userRepo user_repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}
