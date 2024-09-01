package services

import (
	"ayobelajar-app-backend/inputs"
	"ayobelajar-app-backend/models"
	"ayobelajar-app-backend/repositories"
	"strings"
	"time"
)

type CourseCategoryService interface {
	GetCourseCategories() ([]models.CourseCategory, error)
	GetCourseCategory(courseCategoryCode string) (models.CourseCategory, error)
	CreateCourseCategory(createCourseCategoryInput inputs.CreateCourseCategoryInput, currentUser map[string]string) (models.CourseCategory, error)
	UpdateCourseCategory(courseCategoryCode string, updateCourseCategoryInput inputs.UpdateCourseCategoryInput, currentUser map[string]string) (models.CourseCategory, error)
	DeleteCourseCategory(courseCategoryCode string) (models.CourseCategory, error)
}

type courseCategoryService struct {
	courseCategoryRepository repositories.CourseCategoryRepository
}

func NewCourseCategoryService(courseCategoryRepository repositories.CourseCategoryRepository) *courseCategoryService {
	return &courseCategoryService{courseCategoryRepository}
}

func (ccs *courseCategoryService) GetCourseCategories() ([]models.CourseCategory, error) {
	courseCategories, err := ccs.courseCategoryRepository.GetAll()

	return courseCategories, err
}

func (ccs *courseCategoryService) GetCourseCategory(courseCategoryCode string) (models.CourseCategory, error) {
	courseCategory, err := ccs.courseCategoryRepository.GetCourseCategory(courseCategoryCode)

	return courseCategory, err
}

func (ccs *courseCategoryService) CreateCourseCategory(createCourseCategoryInput inputs.CreateCourseCategoryInput, currentUser map[string]string) (models.CourseCategory, error) {
	admin_uuid := currentUser["current_user_uuid"]
	admin_username := currentUser["current_user_username"]

	course_category_to_slug := strings.ToLower(strings.ReplaceAll(createCourseCategoryInput.CourseCategoryName, " ", "-"))

	courseCategory := models.CourseCategory{
		CourseCategoryCode:            course_category_to_slug,
		CourseCategoryName:            createCourseCategoryInput.CourseCategoryName,
		CourseCategoryDescription:     createCourseCategoryInput.CourseCategoryDescription,
		CourseCategoryTags:            createCourseCategoryInput.CourseCategoryTags,
		CourseCategoryStatusCd:        "active",
		CourseCategoryCreatedDate:     time.Now(),
		CourseCategoryCreatedUserUuid: admin_uuid,
		CourseCategoryCreatedUsername: admin_username,
	}

	newCourseCategory, err := ccs.courseCategoryRepository.Create(courseCategory)

	return newCourseCategory, err
}

func (ccs *courseCategoryService) UpdateCourseCategory(courseCategoryCode string, updateCourseCategoryInput inputs.UpdateCourseCategoryInput, currentUser map[string]string) (models.CourseCategory, error) {
	checkCourseCategory, _ := ccs.courseCategoryRepository.GetCourseCategory(courseCategoryCode)

	admin_uuid := currentUser["current_user_uuid"]
	admin_username := currentUser["current_user_username"]

	updated_course_category_to_slug := strings.ToLower(strings.ReplaceAll(updateCourseCategoryInput.CourseCategoryName, " ", "-"))

	checkCourseCategory.CourseCategoryCode = updated_course_category_to_slug
	checkCourseCategory.CourseCategoryName = updateCourseCategoryInput.CourseCategoryName
	checkCourseCategory.CourseCategoryDescription = updateCourseCategoryInput.CourseCategoryDescription
	checkCourseCategory.CourseCategoryTags = updateCourseCategoryInput.CourseCategoryTags
	checkCourseCategory.CourseCategoryUpdatedDate = time.Now()
	checkCourseCategory.CourseCategoryUpdatedUserUuid = admin_uuid
	checkCourseCategory.CourseCategoryUpdatedUsername = admin_username

	updatedCourseCategory, err := ccs.courseCategoryRepository.Update(checkCourseCategory)

	return updatedCourseCategory, err
}

func (ccs *courseCategoryService) DeleteCourseCategory(courseCategoryCode string) (models.CourseCategory, error) {
	checkCourseCategory, _ := ccs.courseCategoryRepository.GetCourseCategory(courseCategoryCode)

	deletedCourseCategory, err := ccs.courseCategoryRepository.Delete(checkCourseCategory)

	return deletedCourseCategory, err
}
