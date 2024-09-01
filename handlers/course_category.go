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

type courseCategoryHandler struct {
	courseCategoryService services.CourseCategoryService
}

func NewCourseCategoryHandler(courseCategoryService services.CourseCategoryService) *courseCategoryHandler {
	return &courseCategoryHandler{courseCategoryService}
}

func (cch *courseCategoryHandler) GetAllCourseCategoriesHandler(c *gin.Context) {
	courseCategories, err := cch.courseCategoryService.GetCourseCategories()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "200",
			"message": "All course categories is empty! Please create a new course categories now!",
		})
		return
	}

	var courseCategoriesRsps []responses.CourseCategoriesResponse

	for _, courseCategory := range courseCategories {
		courseCategoryRsps := responses.ConvertToCourseCategoriesResponse(courseCategory)

		courseCategoriesRsps = append(courseCategoriesRsps, courseCategoryRsps)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "All course categories succesfully appeared!",
		"data":    courseCategoriesRsps,
	})
}

func (cch *courseCategoryHandler) GetOneCourseCategoryHandler(c *gin.Context) {
	courseCategoryCode := c.Param("course_category_code")

	checkCourseCategory, err := cch.courseCategoryService.GetCourseCategory(courseCategoryCode)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "The course category " + courseCategoryCode + " is not appeared! Please try again!",
		})
		return
	}

	courseCategoryRsps := responses.ConvertToCourseCategoriesResponse(checkCourseCategory)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "The course category " + courseCategoryCode + " data is appeared!",
		"data":    courseCategoryRsps,
	})
}

func (cch *courseCategoryHandler) CreateNewCourseCategoryHandler(c *gin.Context) {
	var createNewCourseCategoryInput inputs.CreateCourseCategoryInput

	errCreateNewCourseCategoryInput := c.ShouldBindJSON(&createNewCourseCategoryInput)

	if errCreateNewCourseCategoryInput != nil {
		errorMessages := []string{}

		for _, e := range errCreateNewCourseCategoryInput.(validator.ValidationErrors) {
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

	newCreateCourseCategory, err := cch.courseCategoryService.CreateCourseCategory(createNewCourseCategoryInput, current_user_data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "400",
			"status":  "Process creating course category is failed! Please try again!",
		})
		return
	}

	createCourseCategoryRsps := responses.ConvertToCourseCategoriesResponse(newCreateCourseCategory)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "A new category course is succesfully created!",
		"data":    createCourseCategoryRsps,
	})
}

func (cch *courseCategoryHandler) UpdateOneCourseCategoryHandler(c *gin.Context) {
	courseCategoryCode := c.Param("course_category_code")

	_, err := cch.courseCategoryService.GetCourseCategory(courseCategoryCode)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "The course category " + courseCategoryCode + " is not appeared! Please try again!",
		})
		return
	}

	var updateCourseCategoryInput inputs.UpdateCourseCategoryInput

	errUpdateCourseCategoryInput := c.ShouldBindJSON(&updateCourseCategoryInput)

	if errUpdateCourseCategoryInput != nil {
		errorMessages := []string{}

		for _, e := range errUpdateCourseCategoryInput.(validator.ValidationErrors) {
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

	updateCourseCategory, err := cch.courseCategoryService.UpdateCourseCategory(courseCategoryCode, updateCourseCategoryInput, current_user_data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Process update course category " + courseCategoryCode + " is failed! Please try again!",
		})
		return
	}

	updateCourseCategoryRsps := responses.ConvertToCourseCategoriesResponse(updateCourseCategory)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "The course category " + courseCategoryCode + " data is succesfully updated!",
		"data":    updateCourseCategoryRsps,
	})
}

func (cch *courseCategoryHandler) DeleteOneCourseCategoryHandler(c *gin.Context) {
	courseCategoryCode := c.Param("course_category_code")

	_, err := cch.courseCategoryService.GetCourseCategory(courseCategoryCode)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "The course category " + courseCategoryCode + " is not appeared! Please try again!",
		})
		return
	}

	_, errDeleteCourseCategory := cch.courseCategoryService.DeleteCourseCategory(courseCategoryCode)

	if errDeleteCourseCategory != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Process deleting a course category " + courseCategoryCode + " is failed! Please try again!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "The course category data " + courseCategoryCode + " is succesfully deleted!",
	})

}
