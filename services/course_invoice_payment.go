package services

import (
	"ayobelajar-app-backend/inputs"
	"ayobelajar-app-backend/models"
	"ayobelajar-app-backend/repositories"
	"time"
)

type CourseInvoicePaymentService interface {
	GetCourseInvoicePayments() ([]models.CourseInvoicePayment, error)
	GetOwnCourseInvoicePayments(userUUID string) ([]models.CourseInvoicePayment, error)
	GetCourseInvoicePayment(courseInvoicePaymentCode string) (models.CourseInvoicePayment, error)
	CheckCourseQuantity(courseSlugCode string) (models.Course, error)
	CreateCourseInvoicePayment(courseInvoicePaymentInput inputs.CourseInvoicePaymentInput, currentUser map[string]string) (models.CourseInvoicePayment, error)
	DeleteCourseInvoicePayment(courseInvoicePaymentCode string) (models.CourseInvoicePayment, error)
}

type courseInvoicePaymentService struct {
	courseInvoicePaymentRepository repositories.CourseInvoicePaymentRepository
}

func NewCourseInvoicePaymentService(courseInvoicePaymentRepository repositories.CourseInvoicePaymentRepository) *courseInvoicePaymentService {
	return &courseInvoicePaymentService{courseInvoicePaymentRepository}
}

func (cips *courseInvoicePaymentService) GetCourseInvoicePayments() ([]models.CourseInvoicePayment, error) {
	courseInvoicePayments, err := cips.courseInvoicePaymentRepository.GetAll()

	return courseInvoicePayments, err
}

func (cips *courseInvoicePaymentService) GetOwnCourseInvoicePayments(userUUID string) ([]models.CourseInvoicePayment, error) {
	courseInvoicePayments, err := cips.courseInvoicePaymentRepository.GetOwn(userUUID)

	return courseInvoicePayments, err
}

func (cips *courseInvoicePaymentService) GetCourseInvoicePayment(courseInvoicePaymentCode string) (models.CourseInvoicePayment, error) {
	courseInvoicePayment, err := cips.courseInvoicePaymentRepository.GetCourseInvoicePayment(courseInvoicePaymentCode)

	return courseInvoicePayment, err
}

func (cips *courseInvoicePaymentService) CheckCourseQuantity(courseSlugCode string) (models.Course, error) {
	checkCourse, err := cips.courseInvoicePaymentRepository.GetCourseForPayment(courseSlugCode)

	return checkCourse, err
}

func (cips *courseInvoicePaymentService) CreateCourseInvoicePayment(courseInvoicePaymentInput inputs.CourseInvoicePaymentInput, currentUser map[string]string) (models.CourseInvoicePayment, error) {
	current_user_uuid := currentUser["current_user_uuid"]
	current_username := currentUser["current_user_username"]

	checkCourse, _ := cips.courseInvoicePaymentRepository.GetCourseForPayment(courseInvoicePaymentInput.CourseSlugCode)

	newPrice := courseInvoicePaymentInput.CourseInvoicePaymentQuantity * checkCourse.CoursePrice

	courseInvoicePayment := models.CourseInvoicePayment{
		UserUUID:                            current_user_uuid,
		CourseSlugCode:                      courseInvoicePaymentInput.CourseSlugCode,
		CourseCartCode:                      courseInvoicePaymentInput.CourseCartCode,
		CourseInvoicePaymentCode:            courseInvoicePaymentInput.CourseInvoicePaymentCode,
		CourseInvoicePaymentPrice:           newPrice,
		CourseInvoicePaymentQuantity:        courseInvoicePaymentInput.CourseInvoicePaymentQuantity,
		CourseInvoicePaymentStatusCd:        "pending",
		CourseInvoicePaymentCreatedDate:     time.Now(),
		CourseInvoicePaymentCreatedUserUuid: current_user_uuid,
		CourseInvoicePaymentCreatedUsername: current_username,
	}

	newCourseInvoicePayment, err := cips.courseInvoicePaymentRepository.Create(courseInvoicePayment)

	newQuantity := checkCourse.CourseQuantity - courseInvoicePaymentInput.CourseInvoicePaymentQuantity

	cips.courseInvoicePaymentRepository.UpdateQuantity(checkCourse.CourseSlugCode, newQuantity)

	return newCourseInvoicePayment, err
}

func (cips *courseInvoicePaymentService) DeleteCourseInvoicePayment(courseInvoicePaymentCode string) (models.CourseInvoicePayment, error) {
	checkCourse, _ := cips.courseInvoicePaymentRepository.GetCourseInvoicePayment(courseInvoicePaymentCode)

	deleteCourseInvoicePayment, err := cips.courseInvoicePaymentRepository.Delete(checkCourse)

	return deleteCourseInvoicePayment, err
}
