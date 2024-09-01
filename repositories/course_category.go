package repositories

import (
	"ayobelajar-app-backend/models"

	"gorm.io/gorm"
)

type CourseCategoryRepository interface {
	GetAll() ([]models.CourseCategory, error)
	GetCourseCategory(courseCategoryCode string) (models.CourseCategory, error)
	Create(courseCategory models.CourseCategory) (models.CourseCategory, error)
	Update(courseCategory models.CourseCategory) (models.CourseCategory, error)
	Delete(courseCategory models.CourseCategory) (models.CourseCategory, error)
}

type courseCategoryRepository struct {
	db *gorm.DB
}

func NewCourseCategoryRepository(db *gorm.DB) *courseCategoryRepository {
	return &courseCategoryRepository{db}
}

func (ccr *courseCategoryRepository) GetAll() ([]models.CourseCategory, error) {
	var courseCategories []models.CourseCategory

	err := ccr.db.Find(&courseCategories).Error

	return courseCategories, err
}

func (ccr *courseCategoryRepository) GetCourseCategory(courseCategoryCode string) (models.CourseCategory, error) {
	var courseCategory models.CourseCategory

	err := ccr.db.Where("course_category_code = ?", courseCategoryCode).First(&courseCategory).Error

	return courseCategory, err
}

func (ccr *courseCategoryRepository) Create(courseCategory models.CourseCategory) (models.CourseCategory, error) {
	err := ccr.db.Create(&courseCategory).Error

	return courseCategory, err
}

func (ccr *courseCategoryRepository) Update(courseCategory models.CourseCategory) (models.CourseCategory, error) {
	err := ccr.db.Save(&courseCategory).Error

	return courseCategory, err
}

func (ccr *courseCategoryRepository) Delete(courseCategory models.CourseCategory) (models.CourseCategory, error) {
	err := ccr.db.Delete(&courseCategory).Error

	return courseCategory, err
}
