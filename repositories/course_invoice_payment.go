package repositories

import (
	"ayobelajar-app-backend/models"

	"gorm.io/gorm"
)

type CourseInvoicePaymentRepository interface {
	GetAll() ([]models.CourseInvoicePayment, error)
	GetOwn(userUUID string) ([]models.CourseInvoicePayment, error)
	GetCourseForPayment(courseSlugCode string) (models.Course, error)
	GetCourseInvoicePayment(courseInvoicePayment string) (models.CourseInvoicePayment, error)
	Create(courseInvoicePayment models.CourseInvoicePayment) (models.CourseInvoicePayment, error)
	Update(courseInvoicePayment models.CourseInvoicePayment) (models.CourseInvoicePayment, error)
	UpdateQuantity(courseSlugCode string, newQuantity int) error
	Delete(courseInvoicePayment models.CourseInvoicePayment) (models.CourseInvoicePayment, error)
}

type courseInvoicePaymentRepository struct {
	db *gorm.DB
}

func NewCourseInvoicePaymentRepository(db *gorm.DB) *courseInvoicePaymentRepository {
	return &courseInvoicePaymentRepository{db}
}

func (cipr *courseInvoicePaymentRepository) GetAll() ([]models.CourseInvoicePayment, error) {
	var courseInvoicePayments []models.CourseInvoicePayment

	err := cipr.db.Find(&courseInvoicePayments).Error

	return courseInvoicePayments, err
}

func (cipr *courseInvoicePaymentRepository) GetOwn(userUUID string) ([]models.CourseInvoicePayment, error) {
	var courseInvoicePayments []models.CourseInvoicePayment

	err := cipr.db.Where("user_uuid = ?", userUUID).Find(&courseInvoicePayments).Error

	return courseInvoicePayments, err
}

func (cipr *courseInvoicePaymentRepository) GetCourseForPayment(courseSlugCode string) (models.Course, error) {
	var course models.Course

	err := cipr.db.Where("course_slug_code = ?", courseSlugCode).First(&course).Error

	return course, err
}

func (cipr *courseInvoicePaymentRepository) GetCourseInvoicePayment(courseInvoicePaymentCode string) (models.CourseInvoicePayment, error) {
	var courseInvoicePayment models.CourseInvoicePayment

	err := cipr.db.Where("course_invoice_payment_id = ?", courseInvoicePaymentCode).First(&courseInvoicePayment).Error

	return courseInvoicePayment, err
}

func (cipr *courseInvoicePaymentRepository) Create(courseInvoicePayment models.CourseInvoicePayment) (models.CourseInvoicePayment, error) {
	err := cipr.db.Create(&courseInvoicePayment).Error

	return courseInvoicePayment, err
}

func (cipr *courseInvoicePaymentRepository) Update(courseInvoicePayment models.CourseInvoicePayment) (models.CourseInvoicePayment, error) {
	err := cipr.db.Save(&courseInvoicePayment).Error

	return courseInvoicePayment, err
}

func (cipr *courseInvoicePaymentRepository) UpdateQuantity(courseSlugCode string, newQuantity int) error {
	err := cipr.db.Model(&models.Course{}).Where("course_slug_code = ?", courseSlugCode).Update("course_quantity", newQuantity).Error

	return err
}

func (cipr *courseInvoicePaymentRepository) Delete(courseInvoicePayment models.CourseInvoicePayment) (models.CourseInvoicePayment, error) {
	err := cipr.db.Delete(&courseInvoicePayment).Error

	return courseInvoicePayment, err
}
