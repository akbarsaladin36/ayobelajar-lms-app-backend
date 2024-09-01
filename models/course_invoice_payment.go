package models

import "time"

type CourseInvoicePayment struct {
	CourseInvoicePaymentId              int    `json:"course_invoice_payment_id" gorm:"primaryKey"`
	UserUUID                            string `json:"user_uuid" gorm:"type:varchar(200)"`
	CourseSlugCode                      string `json:"course_slug_code" gorm:"type:varchar(100)"`
	CourseCartCode                      string `json:"course_cart_code" gorm:"type:varchar(100)"`
	CourseInvoicePaymentCode            string `json:"course_invoice_payment_code" gorm:"type:varchar(150)"`
	CourseInvoicePaymentPrice           int    `json:"course_invoice_payment_price" gorm:"type:varchar(100)"`
	CourseInvoicePaymentQuantity        int    `json:"course_invoice_payment_quantity" gorm:"type:varchar(100)"`
	CourseInvoicePaymentStatusCd        string `json:"course_invoice_payment_status_cd" gorm:"type:varchar(50)"`
	CourseInvoicePaymentCreatedDate     time.Time
	CourseInvoicePaymentCreatedUserUuid string `json:"course_invoice_payment_created_user_uuid" gorm:"type:varchar(200)"`
	CourseInvoicePaymentCreatedUsername string `json:"course_invoice_payment_created_username" gorm:"type:varchar(150)"`
	CourseInvoicePaymentUpdatedDate     time.Time
	CourseInvoicePaymentUpdatedUserUuid string `json:"course_invoice_payment_updated_user_uuid" gorm:"type:varchar(200)"`
	CourseInvoicePaymentUpdatedUsername string `json:"course_invoice_payment_updated_username" gorm:"type:varchar(150)"`
}
