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

type courseCartHandler struct {
	courseCartService services.CourseCartService
}

func NewCourseCartHandler(courseCartService services.CourseCartService) *courseCartHandler {
	return &courseCartHandler{courseCartService}
}

func (ccth *courseCartHandler) GetAllCourseCartsHandler(c *gin.Context) {
	courseCarts, err := ccth.courseCartService.GetCourseCarts()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "200",
			"message": "All course cart is empty!",
		})
		return
	}

	var courseCartsResponse []responses.CourseCartResponse

	for _, courseCart := range courseCarts {
		courseCartResponse := responses.ConvertToCourseCartResponse(courseCart)

		courseCartsResponse = append(courseCartsResponse, courseCartResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "All course carts is succesfully appeared!",
		"data":    courseCartsResponse,
	})
}

func (ccth *courseCartHandler) GetMyCourseCartsHandler(c *gin.Context) {
	userUUID, _, _, _ := middleware.CurrentUser(c)

	courseCarts, err := ccth.courseCartService.GetOwnCourseCarts(userUUID)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "200",
			"message": "My course carts is empty!",
		})
		return
	}

	var courseCartsResponse []responses.CourseCartResponse

	for _, courseCart := range courseCarts {
		courseCartResponse := responses.ConvertToCourseCartResponse(courseCart)

		courseCartsResponse = append(courseCartsResponse, courseCartResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "All course carts is succesfully appeared!",
		"data":    courseCartsResponse,
	})
}

func (ccth *courseCartHandler) GetCourseCartHandler(c *gin.Context) {
	courseCartCode := c.Param("course_cart_code")

	courseCart, err := ccth.courseCartService.GetCourseCart(courseCartCode)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "The course cart " + courseCartCode + " is not appeared! Please try again!",
		})
		return
	}

	courseCartRsps := responses.ConvertToCourseCartResponse(courseCart)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "The course cart " + courseCartCode + " is succesfully appeared!",
		"data":    courseCartRsps,
	})
}

func (ccth *courseCartHandler) CreateNewCourseCartHandler(c *gin.Context) {
	var createNewCourseCartInput inputs.CreateCourseCartInput

	errCreateNewCourseCartInput := c.ShouldBindJSON(&createNewCourseCartInput)

	if errCreateNewCourseCartInput != nil {
		errorMessages := []string{}

		for _, e := range errCreateNewCourseCartInput.(validator.ValidationErrors) {
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

	newCourseCart, err := ccth.courseCartService.CreateCourseCart(createNewCourseCartInput, current_user_data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Process creating course cart is failed! Please try again!",
		})
		return
	}

	createCourseCartRsps := responses.ConvertToCourseCartResponse(newCourseCart)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "A new course cart is succesfully created!",
		"data":    createCourseCartRsps,
	})
}

func (ccth *courseCartHandler) UpdateOneCourseCartHandler(c *gin.Context) {
	courseCartCode := c.Param("course_cart_code")

	_, err := ccth.courseCartService.GetCourseCart(courseCartCode)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "The course cart " + courseCartCode + " is not appeared! Please try again!",
		})
		return
	}

	var updateCourseCartInput inputs.UpdateCourseCartInput

	errUpdateCourseCartInput := c.ShouldBindJSON(&updateCourseCartInput)

	if errUpdateCourseCartInput != nil {
		errorMessages := []string{}

		for _, e := range errUpdateCourseCartInput.(validator.ValidationErrors) {
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

	updateCourseCart, err := ccth.courseCartService.UpdateCourseCart(courseCartCode, updateCourseCartInput, current_user_data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Process updating course cart " + courseCartCode + " data is failed! Please try again!",
		})
		return
	}

	updateCourseCartRsps := responses.ConvertToCourseCartResponse(updateCourseCart)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "The course cart " + courseCartCode + " data is succesfully updated!",
		"data":    updateCourseCartRsps,
	})
}

func (ccth *courseCartHandler) DeleteOneCourseCartHandler(c *gin.Context) {
	courseCartCode := c.Param("course_cart_code")

	_, err := ccth.courseCartService.GetCourseCart(courseCartCode)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "The course cart " + courseCartCode + " is not appeared! Please try again!",
		})
		return
	}

	_, errDeleteCourseCart := ccth.courseCartService.DeleteCourseCart(courseCartCode)

	if errDeleteCourseCart != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Process deleting course cart " + courseCartCode + " data are failed! Please try again!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "The course cart " + courseCartCode + " data are succesfully deleted!",
	})
}
