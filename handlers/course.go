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

type courseHandler struct {
	courseService services.CourseService
}

func NewCourseHandler(courseService services.CourseService) *courseHandler {
	return &courseHandler{courseService}
}

func (ch *courseHandler) GetAllCoursesHandler(c *gin.Context) {
	courses, err := ch.courseService.GetCourses()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "The courses data is empty! Please create a new course now!",
		})
		return
	}

	var coursesResponse []responses.CoursesResponse

	for _, course := range courses {
		courseResponse := responses.ConvertToCoursesResponse(course)

		coursesResponse = append(coursesResponse, courseResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "All courses data is succesfully appeared!",
		"data":    coursesResponse,
	})
}

func (ch *courseHandler) GetOneCourseHandler(c *gin.Context) {
	courseSlugCode := c.Param("course_slug_code")

	course, err := ch.courseService.GetCourse(courseSlugCode)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "The " + courseSlugCode + " data is not appeared! Please try again!",
		})
		return
	}

	courseResponse := responses.ConvertToCoursesResponse(course)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "The " + courseSlugCode + " data is not appeared! Please try again!",
		"data":    courseResponse,
	})
}

func (ch *courseHandler) CreateNewCourseHandler(c *gin.Context) {
	var createCourseInput inputs.CreateCourseInput

	errCreateCourseInput := c.ShouldBindJSON(&createCourseInput)

	if errCreateCourseInput != nil {
		errorMessages := []string{}

		for _, e := range errCreateCourseInput.(validator.ValidationErrors) {
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

	newCourse, err := ch.courseService.CreateCourse(createCourseInput, current_user_data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Process creating a course data is failed! Please try again!",
		})
		return
	}

	createCourseRsps := responses.ConvertToCreateCourseResponse(newCourse)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "A new course data is succesfully created!",
		"data":    createCourseRsps,
	})

}

func (ch *courseHandler) UpdateOneCourseHandler(c *gin.Context) {
	courseSlugCode := c.Param("course_slug_code")

	var updateCourseInput inputs.UpdateCourseInput

	errUpdateCourseInput := c.ShouldBindJSON(&updateCourseInput)

	if errUpdateCourseInput != nil {
		errorMessages := []string{}

		for _, e := range errUpdateCourseInput.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": errorMessages,
		})

		return
	}

	_, err := ch.courseService.GetCourse(courseSlugCode)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "The data for course " + courseSlugCode + " is not appeared! Please try again!",
		})
		return
	}

	userUUID, userName, _, _ := middleware.CurrentUser(c)

	current_user_data := map[string]string{
		"current_user_uuid":     userUUID,
		"current_user_username": userName,
	}

	updateCourse, errUpdateCourse := ch.courseService.UpdateCourse(courseSlugCode, updateCourseInput, current_user_data)

	if errUpdateCourse != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Process updating data for course " + courseSlugCode + " is failed! Please try again!",
		})
		return
	}

	updateCourseRsps := responses.ConvertToUpdateCourseResponse(updateCourse)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Updating data for course " + courseSlugCode + " is succesfully updated!",
		"data":    updateCourseRsps,
	})
}

func (ch *courseHandler) DeleteOneCourseHandler(c *gin.Context) {
	courseSlugCode := c.Param("course_slug_code")

	_, err := ch.courseService.GetCourse(courseSlugCode)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "The data for course " + courseSlugCode + " is not appeared! Please try again!",
		})
		return
	}

	_, errDeleteCourse := ch.courseService.DeleteCourse(courseSlugCode)

	if errDeleteCourse != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Process delete course " + courseSlugCode + " is failed! Please try again!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "A course " + courseSlugCode + " data is succesfully deleted!",
	})
}
