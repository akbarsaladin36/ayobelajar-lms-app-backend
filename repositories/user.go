package repositories

import (
	"ayobelajar-app-backend/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() ([]models.User, error)
	GetUser(username string) (models.User, error)
	Create(user models.User) (models.User, error)
	Update(user models.User) (models.User, error)
	Delete(user models.User) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetAll() ([]models.User, error) {
	var users []models.User

	err := ur.db.Find(&users).Error

	return users, err
}

func (ur *userRepository) GetUser(username string) (models.User, error) {
	var user models.User

	err := ur.db.Where("user_username = ?", username).First(&user).Error

	return user, err
}

func (ur *userRepository) Create(user models.User) (models.User, error) {
	err := ur.db.Create(&user).Error

	return user, err
}

func (ur *userRepository) Update(user models.User) (models.User, error) {
	err := ur.db.Save(&user).Error

	return user, err
}

func (ur *userRepository) Delete(user models.User) (models.User, error) {
	err := ur.db.Delete(&user).Error

	return user, err
}
