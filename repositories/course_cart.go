package repositories

import (
	"ayobelajar-app-backend/models"

	"gorm.io/gorm"
)

type CourseCartRepository interface {
	GetAll() ([]models.CourseCart, error)
	GetAllByUserUuid(userUUID string) ([]models.CourseCart, error)
	GetCourseCart(courseCartCode string) (models.CourseCart, error)
	Create(courseCart models.CourseCart) (models.CourseCart, error)
	Update(courseCart models.CourseCart) (models.CourseCart, error)
	Delete(courseCart models.CourseCart) (models.CourseCart, error)
}

type courseCartRepository struct {
	db *gorm.DB
}

func NewCourseCartRepository(db *gorm.DB) *courseCartRepository {
	return &courseCartRepository{db}
}

func (cctr *courseCartRepository) GetAll() ([]models.CourseCart, error) {
	var courseCarts []models.CourseCart

	err := cctr.db.Find(&courseCarts).Error

	return courseCarts, err
}

func (cctr *courseCartRepository) GetAllByUserUuid(userUUID string) ([]models.CourseCart, error) {
	var courseCarts []models.CourseCart

	err := cctr.db.Where("user_uuid = ?", userUUID).Find(&courseCarts).Error

	return courseCarts, err
}

func (cctr *courseCartRepository) GetCourseCart(courseCartCode string) (models.CourseCart, error) {
	var courseCart models.CourseCart

	err := cctr.db.Where("course_cart_code = ?", courseCartCode).First(&courseCart).Error

	return courseCart, err
}

func (cctr *courseCartRepository) Create(courseCart models.CourseCart) (models.CourseCart, error) {
	err := cctr.db.Create(&courseCart).Error

	return courseCart, err
}

func (cctr *courseCartRepository) Update(courseCart models.CourseCart) (models.CourseCart, error) {
	err := cctr.db.Save(&courseCart).Error

	return courseCart, err
}

func (cctr *courseCartRepository) Delete(courseCart models.CourseCart) (models.CourseCart, error) {
	err := cctr.db.Delete(&courseCart).Error

	return courseCart, err
}
