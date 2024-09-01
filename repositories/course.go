package repositories

import (
	"ayobelajar-app-backend/models"

	"gorm.io/gorm"
)

type CourseRepository interface {
	GetAll() ([]models.Course, error)
	GetCourse(courseSlugCode string) (models.Course, error)
	Create(course models.Course) (models.Course, error)
	Update(course models.Course) (models.Course, error)
	Delete(course models.Course) (models.Course, error)
}

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) *courseRepository {
	return &courseRepository{db}
}

func (cr *courseRepository) GetAll() ([]models.Course, error) {
	var courses []models.Course

	err := cr.db.Find(&courses).Error

	return courses, err
}

func (cr *courseRepository) GetCourse(courseSlugCode string) (models.Course, error) {
	var course models.Course

	err := cr.db.Where("course_slug_code = ?", courseSlugCode).First(&course).Error

	return course, err
}

func (cr *courseRepository) Create(course models.Course) (models.Course, error) {
	err := cr.db.Create(&course).Error

	return course, err
}

func (cr *courseRepository) Update(course models.Course) (models.Course, error) {
	err := cr.db.Save(&course).Error

	return course, err
}

func (cr *courseRepository) Delete(course models.Course) (models.Course, error) {
	err := cr.db.Delete(&course).Error

	return course, err
}
