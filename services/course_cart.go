package services

import (
	"ayobelajar-app-backend/inputs"
	"ayobelajar-app-backend/models"
	"ayobelajar-app-backend/repositories"
	"time"
)

type CourseCartService interface {
	GetCourseCarts() ([]models.CourseCart, error)
	GetOwnCourseCarts(userUUID string) ([]models.CourseCart, error)
	GetCourseCart(courseCartCode string) (models.CourseCart, error)
	CreateCourseCart(createCourseCartInput inputs.CreateCourseCartInput, currentUser map[string]string) (models.CourseCart, error)
	UpdateCourseCart(courseCartCode string, updateCourseCartInput inputs.UpdateCourseCartInput, currentUser map[string]string) (models.CourseCart, error)
	DeleteCourseCart(courseCartCode string) (models.CourseCart, error)
}

type courseCartService struct {
	courseCartRepository repositories.CourseCartRepository
}

func NewCourseCartService(courseCartRepository repositories.CourseCartRepository) *courseCartService {
	return &courseCartService{courseCartRepository}
}

func (ccts *courseCartService) GetCourseCarts() ([]models.CourseCart, error) {
	courseCarts, err := ccts.courseCartRepository.GetAll()

	return courseCarts, err
}

func (ccts *courseCartService) GetOwnCourseCarts(userUUID string) ([]models.CourseCart, error) {
	courseCarts, err := ccts.courseCartRepository.GetAllByUserUuid(userUUID)

	return courseCarts, err
}

func (ccts *courseCartService) GetCourseCart(courseCartCode string) (models.CourseCart, error) {
	courseCart, err := ccts.courseCartRepository.GetCourseCart(courseCartCode)

	return courseCart, err
}

func (ccts *courseCartService) CreateCourseCart(createCourseCartInput inputs.CreateCourseCartInput, currentUser map[string]string) (models.CourseCart, error) {
	user_uuid := currentUser["current_user_uuid"]
	user_username := currentUser["current_user_username"]

	courseCart := models.CourseCart{
		UserUuid:                  user_uuid,
		CourseSlugCode:            createCourseCartInput.CourseSlugCode,
		CourseCartCode:            createCourseCartInput.CourseCartCode,
		CourseCartQuantity:        createCourseCartInput.CourseCartQuantity,
		CourseCartCreatedDate:     time.Now(),
		CourseCartCreatedUserUuid: user_uuid,
		CourseCartCreatedUsername: user_username,
	}

	newCourseCart, err := ccts.courseCartRepository.Create(courseCart)

	return newCourseCart, err
}

func (ccts *courseCartService) UpdateCourseCart(courseCartCode string, updateCourseCartInput inputs.UpdateCourseCartInput, currentUser map[string]string) (models.CourseCart, error) {
	checkCourseCart, _ := ccts.courseCartRepository.GetCourseCart(courseCartCode)

	user_uuid := currentUser["current_user_uuid"]
	user_username := currentUser["current_user_username"]

	checkCourseCart.CourseSlugCode = updateCourseCartInput.CourseSlugCode
	checkCourseCart.CourseCartCode = updateCourseCartInput.CourseCartCode
	checkCourseCart.CourseCartQuantity = updateCourseCartInput.CourseCartQuantity
	checkCourseCart.CourseCartUpdatedDate = time.Now()
	checkCourseCart.CourseCartUpdatedUserUuid = user_uuid
	checkCourseCart.CourseCartUpdatedUsername = user_username

	updateCourseCart, err := ccts.courseCartRepository.Update(checkCourseCart)

	return updateCourseCart, err
}

func (ccts *courseCartService) DeleteCourseCart(courseCartCode string) (models.CourseCart, error) {
	checkCourseCart, _ := ccts.courseCartRepository.GetCourseCart(courseCartCode)

	deletedCourseCart, err := ccts.courseCartRepository.Delete(checkCourseCart)

	return deletedCourseCart, err
}
