package services

import (
	"auth-sederhana-go-fiber/dtos"
	"auth-sederhana-go-fiber/repositories"
	"auth-sederhana-go-fiber/utilities"
	errorUtils "auth-sederhana-go-fiber/utilities/error"
	"auth-sederhana-go-fiber/validation"
	"context"
)

type (
	UserService interface {
		Register(ctx context.Context, data dtos.UserRegisterRequest) (dtos.UserRegisterResponse, errorUtils.CustomError)
		Login(ctx context.Context, data dtos.UserLoginRequest) (dtos.UserLoginResponse, errorUtils.CustomError)
		GetById(ctx context.Context, data dtos.UserGetByIdRequest) (dtos.UserGetByIdResponse, errorUtils.CustomError)
	}

	userService struct {
		userRepo repositories.UserRepository
		jwtUtils utilities.JWTUtils
	}
)

func NewUserService(userRepo repositories.UserRepository, jwtUtils utilities.JWTUtils) UserService {
	return &userService{
		userRepo: userRepo,
		jwtUtils: jwtUtils,
	}
}

func (s *userService) Register(ctx context.Context, data dtos.UserRegisterRequest) (dtos.UserRegisterResponse, errorUtils.CustomError) {
	if err := validation.Validate(data); err != nil {
		return dtos.UserRegisterResponse{}, err
	}

	userExist, err := s.userRepo.FindByEmail(ctx, nil, data.Email)
	if err != nil {
		return dtos.UserRegisterResponse{}, errorUtils.ErrInternalServer
	}

	if userExist.Id != 0 {
		return dtos.UserRegisterResponse{}, errorUtils.ErrEmailAlreadyExist
	}

	user, err := s.userRepo.Create(ctx, nil, data)
	if err != nil {
		return dtos.UserRegisterResponse{}, errorUtils.ErrInternalServer
	}

	res := dtos.UserRegisterResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}

	return res, nil
}

func (s *userService) Login(ctx context.Context, data dtos.UserLoginRequest) (dtos.UserLoginResponse, errorUtils.CustomError) {
	if err := validation.Validate(data); err != nil {
		return dtos.UserLoginResponse{}, err
	}

	user, err := s.userRepo.FindByEmail(ctx, nil, data.Email)
	if err != nil {
		return dtos.UserLoginResponse{}, errorUtils.ErrInternalServer
	}

	if user.Id == 0 {
		return dtos.UserLoginResponse{}, errorUtils.ErrWrongEmailOrPassword
	}

	if isPasswordMatch := utilities.CheckPassword(user.Password, data.Password); !isPasswordMatch {
		return dtos.UserLoginResponse{}, errorUtils.ErrWrongEmailOrPassword
	}

	AccessToken, err := s.jwtUtils.GenerateToken(user.Id)

	if err != nil {
		return dtos.UserLoginResponse{}, errorUtils.ErrInternalServer
	}

	return dtos.UserLoginResponse{
		AccessToken: AccessToken,
	}, nil
}

func (s *userService) GetById(ctx context.Context, data dtos.UserGetByIdRequest) (dtos.UserGetByIdResponse, errorUtils.CustomError) {
	if err := validation.Validate(data); err != nil {
		return dtos.UserGetByIdResponse{}, err
	}

	user, err := s.userRepo.FindById(ctx, nil, data.Id)
	if err != nil {
		return dtos.UserGetByIdResponse{}, errorUtils.ErrInternalServer
	}

	if user.Id == 0 {
		return dtos.UserGetByIdResponse{}, errorUtils.ErrUserNotFound
	}

	res := dtos.UserGetByIdResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}

	return res, nil
}
