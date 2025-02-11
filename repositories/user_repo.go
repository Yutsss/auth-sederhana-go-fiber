package repositories

import (
	"auth-sederhana-go-fiber/dtos"
	"auth-sederhana-go-fiber/entities"
	errorUtils "auth-sederhana-go-fiber/utilities/error"
	"context"
	"errors"
	"gorm.io/gorm"
)

type (
	UserRepository interface {
		Create(ctx context.Context, tx *gorm.DB, data dtos.UserRegisterRequest) (entities.User, errorUtils.CustomError)
		FindByEmail(ctx context.Context, tx *gorm.DB, email string) (entities.User, errorUtils.CustomError)
		FindById(ctx context.Context, tx *gorm.DB, id int64) (entities.User, errorUtils.CustomError)
		FindAll(ctx context.Context, tx *gorm.DB) ([]entities.User, errorUtils.CustomError)
	}

	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(ctx context.Context, tx *gorm.DB, data dtos.UserRegisterRequest) (entities.User, errorUtils.CustomError) {
	if tx == nil {
		tx = r.db
	}

	user := entities.User{
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
	}

	err := tx.WithContext(ctx).Create(&user).Error

	if err != nil {
		return entities.User{}, errorUtils.ErrInternalServer
	}

	return user, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, tx *gorm.DB, email string) (entities.User, errorUtils.CustomError) {
	if tx == nil {
		tx = r.db
	}

	var user entities.User

	err := tx.WithContext(ctx).Where("email = ?", email).First(&user).Error

	if err != nil {
		if errors.As(err, &gorm.ErrRecordNotFound) {
			return entities.User{}, nil
		}
		return entities.User{}, errorUtils.ErrInternalServer
	}

	return user, nil
}

func (r *userRepository) FindById(ctx context.Context, tx *gorm.DB, id int64) (entities.User, errorUtils.CustomError) {
	if tx == nil {
		tx = r.db
	}

	var user entities.User

	err := tx.WithContext(ctx).Where("id = ?", id).First(&user).Error

	if err != nil {
		if errors.As(err, &gorm.ErrRecordNotFound) {
			return entities.User{}, nil
		}
		return entities.User{}, errorUtils.ErrInternalServer
	}

	return user, nil
}

func (r *userRepository) FindAll(ctx context.Context, tx *gorm.DB) ([]entities.User, errorUtils.CustomError) {
	if tx == nil {
		tx = r.db
	}

	var users []entities.User

	err := tx.WithContext(ctx).Where("role = ?", "user").Find(&users).Error

	if err != nil {
		if errors.As(err, &gorm.ErrRecordNotFound) {
			return []entities.User{}, nil
		}
		return []entities.User{}, errorUtils.ErrInternalServer
	}

	return users, nil
}
