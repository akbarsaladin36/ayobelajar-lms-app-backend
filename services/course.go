package services

import (
	"ayobelajar-app-backend/inputs"
	"ayobelajar-app-backend/models"
	"ayobelajar-app-backend/repositories"
	"strings"
	"time"
)

type CourseService interface {
	GetCourses() ([]models.Course, error)
	GetCourse(courseSlugCode string) (models.Course, error)
	CreateCourse(createCourseInput inputs.CreateCourseInput, currentUser map[string]string) (models.Course, error)
	UpdateCourse(courseSlugCode string, updateCourseInput inputs.UpdateCourseInput, currentUser map[string]string) (models.Course, error)
	DeleteCourse(courseSlugCode string) (models.Course, error)
}

type courseService struct {
	courseRepository repositories.CourseRepository
}

func NewCourseService(courseRepository repositories.CourseRepository) *courseService {
	return &courseService{courseRepository}
}

func (cs *courseService) GetCourses() ([]models.Course, error) {
	courses, err := cs.courseRepository.GetAll()

	return courses, err
}

func (cs *courseService) GetCourse(courseSlugCode string) (models.Course, error) {
	course, err := cs.courseRepository.GetCourse(courseSlugCode)

	return course, err
}

func (cs *courseService) CreateCourse(createCourseInput inputs.CreateCourseInput, currentUser map[string]string) (models.Course, error) {
	admin_uuid := currentUser["current_user_uuid"]
	admin_username := currentUser["current_user_username"]

	course_to_slug := strings.ToLower(strings.ReplaceAll(createCourseInput.CourseName, " ", "-"))

	course := models.Course{
		CourseCategoryCode:     createCourseInput.CourseCategoryCode,
		CourseSlugCode:         course_to_slug,
		CourseName:             createCourseInput.CourseName,
		CourseTags:             createCourseInput.CourseTags,
		CoursePrice:            createCourseInput.CoursePrice,
		CourseQuantity:         createCourseInput.CourseQuantity,
		CourseStatusCd:         "active",
		CourseCreatedDate:      time.Now(),
		CourserCreatedUserUuid: admin_uuid,
		CourserCreatedUsername: admin_username,
	}

	newCourse, err := cs.courseRepository.Create(course)

	return newCourse, err
}

func (cs *courseService) UpdateCourse(courseSlugCode string, updateCourseInput inputs.UpdateCourseInput, currentUser map[string]string) (models.Course, error) {
	checkCourse, _ := cs.courseRepository.GetCourse(courseSlugCode)

	admin_uuid := currentUser["current_user_uuid"]
	admin_username := currentUser["current_user_username"]

	updated_course_slug_code := strings.ToLower(strings.ReplaceAll(updateCourseInput.CourseName, " ", "-"))

	checkCourse.CourseCategoryCode = updateCourseInput.CourseCategoryCode
	checkCourse.CourseSlugCode = updated_course_slug_code
	checkCourse.CourseName = updateCourseInput.CourseName
	checkCourse.CourseTags = updateCourseInput.CourseTags
	checkCourse.CoursePrice = updateCourseInput.CoursePrice
	checkCourse.CourseQuantity += updateCourseInput.CourseQuantity
	checkCourse.CourseUpdatedDate = time.Now()
	checkCourse.CourseUpdatedUserUuid = admin_uuid
	checkCourse.CourseUpdatedUsername = admin_username

	updateCourse, err := cs.courseRepository.Update(checkCourse)

	return updateCourse, err
}

func (cs *courseService) DeleteCourse(courseSlugCode string) (models.Course, error) {
	checkCourse, _ := cs.courseRepository.GetCourse(courseSlugCode)

	deleteCourse, err := cs.courseRepository.Delete(checkCourse)

	return deleteCourse, err
}
