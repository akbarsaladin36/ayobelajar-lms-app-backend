package responses

import "ayobelajar-app-backend/models"

type CourseInvoicePaymentResponse struct {
	UserUUID                     string `json:"user_uuid"`
	CourseSlugCode               string `json:"course_slug_code"`
	CourseCartCode               string `json:"course_cart_code"`
	CourseInvoicePaymentPrice    int    `json:"course_invoice_payment_price"`
	CourseInvoicePaymentQuantity int    `json:"course_invoice_payment_quantity"`
	CourseInvoicePaymentCode     string `json:"course_invoice_payment_code"`
	CourseInvoicePaymentStatusCd string `json:"course_invoice_payment_status_cd"`
}

func ConvertToCourseInvoicePaymentResponse(courseInvoicePaymentRsps models.CourseInvoicePayment) CourseInvoicePaymentResponse {
	return CourseInvoicePaymentResponse{
		UserUUID:                     courseInvoicePaymentRsps.UserUUID,
		CourseSlugCode:               courseInvoicePaymentRsps.CourseSlugCode,
		CourseCartCode:               courseInvoicePaymentRsps.CourseCartCode,
		CourseInvoicePaymentPrice:    courseInvoicePaymentRsps.CourseInvoicePaymentPrice,
		CourseInvoicePaymentQuantity: courseInvoicePaymentRsps.CourseInvoicePaymentQuantity,
		CourseInvoicePaymentCode:     courseInvoicePaymentRsps.CourseInvoicePaymentCode,
		CourseInvoicePaymentStatusCd: courseInvoicePaymentRsps.CourseInvoicePaymentStatusCd,
	}
}
