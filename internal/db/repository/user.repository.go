package repository

import (
	"context"

	"github.com/Abrahamthefirst/finecore-practice/internal/db/models"
	"github.com/Abrahamthefirst/finecore-practice/internal/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) FindByID(ctx context.Context, id uint) (*entities.User, error) {
	var user models.UserModel
	err := DBFromCtx(ctx, r.db).Preload("Wallets").Where("id = ?", id).First(&user).Error

	if err != nil {
		return nil, err
	}
	return user.ToDomain(), err
}


func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	var user models.UserModel
	err := DBFromCtx(ctx, r.db).Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}
	return user.ToDomain(), err
}

func (r *UserRepository) FindAll(ctx context.Context, email string) (*[]entities.User, error) {
	var users []models.UserModel
	err := DBFromCtx(ctx, r.db).Find(&users).Error

	if err != nil {
		return nil, err
	}

	var usersList []entities.User

	for _, user := range users {
		usersList = append(usersList, *user.ToDomain())
	}

	return &usersList, err
}

func (r *UserRepository) Create(ctx context.Context, input *entities.User) (*entities.User, error) {

	user := models.UserModel{FirstName: &input.Firstname, LastName: &input.LastName, Username: input.Username, Email: input.Email, Password: input.Password, AvatarKey: &input.AvatarKey}

	err := DBFromCtx(ctx, r.db).Create(&user).Error

	if err != nil {
		return nil, err
	}
	return user.ToDomain(), err
}
