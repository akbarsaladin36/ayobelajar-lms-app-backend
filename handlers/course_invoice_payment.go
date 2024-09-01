package handlers

import (
	"ayobelajar-app-backend/inputs"
	"ayobelajar-app-backend/middleware"
	"ayobelajar-app-backend/responses"
	"ayobelajar-app-backend/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type courseInvoicePaymentHandler struct {
	courseInvoicePaymentService services.CourseInvoicePaymentService
}

func NewCourseInvoicePaymentHandler(courseInvoicePaymentService services.CourseInvoicePaymentService) *courseInvoicePaymentHandler {
	return &courseInvoicePaymentHandler{courseInvoicePaymentService}
}

func (ciph *courseInvoicePaymentHandler) GetAllCourseInvoicePaymentsHandler(c *gin.Context) {
	courseInvoicePayments, err := ciph.courseInvoicePaymentService.GetCourseInvoicePayments()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "All course invoice payments data is empty!",
		})
		return
	}

	var courseInvoicePaymentsResponse []responses.CourseInvoicePaymentResponse

	for _, courseInvoicePayment := range courseInvoicePayments {
		courseInvoicePaymentResponse := responses.ConvertToCourseInvoicePaymentResponse(courseInvoicePayment)

		courseInvoicePaymentsResponse = append(courseInvoicePaymentsResponse, courseInvoicePaymentResponse)
	}

	if courseInvoicePaymentsResponse == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "200",
			"message": "All course invoice payments data is empty!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "All course invoice payments data is succesfully appeared!",
		"data":    courseInvoicePaymentsResponse,
	})

}

func (ciph *courseInvoicePaymentHandler) GetMyCourseInvoicePaymentsHandler(c *gin.Context) {
	user_uuid, _, _, _ := middleware.CurrentUser(c)

	courseInvoicePayments, err := ciph.courseInvoicePaymentService.GetOwnCourseInvoicePayments(user_uuid)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "All course invoice payments data is empty!",
		})
		return
	}

	var courseInvoicePaymentsResponse []responses.CourseInvoicePaymentResponse

	for _, courseInvoicePayment := range courseInvoicePayments {
		courseInvoicePaymentResponse := responses.ConvertToCourseInvoicePaymentResponse(courseInvoicePayment)

		courseInvoicePaymentsResponse = append(courseInvoicePaymentsResponse, courseInvoicePaymentResponse)
	}

	if courseInvoicePaymentsResponse == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "200",
			"message": "All course invoice payments data is empty!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "All my course invoice payments data is succesfully appeared!",
		"data":    courseInvoicePaymentsResponse,
	})

}

func (ciph *courseInvoicePaymentHandler) GetOneCourseInvoicePaymentHandler(c *gin.Context) {
	courseInvoicePaymentCode := c.Param("course_invoice_payment_code")

	checkCourseInvoicePayment, err := ciph.courseInvoicePaymentService.GetCourseInvoicePayment(courseInvoicePaymentCode)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "The course invoice payment data for " + courseInvoicePaymentCode + "is not appeared!",
		})
		return
	}

	courseInvoicePaymentRsps := responses.ConvertToCourseInvoicePaymentResponse(checkCourseInvoicePayment)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "The course invoice payment data for " + courseInvoicePaymentCode + "is successfully appeared!",
		"data":    courseInvoicePaymentRsps,
	})

}

func (ciph *courseInvoicePaymentHandler) CreateNewCourseInvoicePaymentHandler(c *gin.Context) {
	var createCourseInvoicePaymentInput inputs.CourseInvoicePaymentInput

	errCreateCourseInvoicePaymentInput := c.ShouldBindJSON(&createCourseInvoicePaymentInput)

	if errCreateCourseInvoicePaymentInput != nil {
		errorMessages := []string{}

		for _, e := range errCreateCourseInvoicePaymentInput.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": errorMessages,
		})

		return
	}

	userUUID, userName, _, _ := middleware.CurrentUser(c)

	current_user_data := map[string]string{
		"current_user_uuid":     userUUID,
		"current_user_username": userName,
	}

	checkQuantityCourseInvoicePayment, _ := ciph.courseInvoicePaymentService.CheckCourseQuantity(createCourseInvoicePaymentInput.CourseSlugCode)

	if checkQuantityCourseInvoicePayment.CourseQuantity < createCourseInvoicePaymentInput.CourseInvoicePaymentQuantity {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Insufficient course quantity! Please ask admin to adding more course quantity!",
		})
		return
	}

	newCourseInvoicePayment, err := ciph.courseInvoicePaymentService.CreateCourseInvoicePayment(createCourseInvoicePaymentInput, current_user_data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Process creating invoice for payment course is failed! Please try again!",
		})
		return
	}

	newCourseInvoicePaymentRsps := responses.ConvertToCourseInvoicePaymentResponse(newCourseInvoicePayment)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "A new invoice payment for course is succesfully created!",
		"data":    newCourseInvoicePaymentRsps,
	})
}
