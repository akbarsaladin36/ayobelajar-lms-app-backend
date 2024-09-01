package inputs

type CourseInvoicePaymentInput struct {
	UserUUID                     string `json:"user_uuid"`
	CourseSlugCode               string `json:"course_slug_code"`
	CourseCartCode               string `json:"course_cart_code"`
	CourseInvoicePaymentQuantity int    `json:"course_invoice_payment_quantity"`
	CourseInvoicePaymentCode     string `json:"course_invoice_payment_code"`
	CourseInvoicePaymentStatusCd string `json:"course_invoice_payment_status_cd"`
}
