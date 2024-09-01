package inputs

type CreateCourseInput struct {
	CourseCategoryCode string `json:"course_category_code"`
	CourseName         string `json:"course_name"`
	CourseTags         string `json:"course_tags"`
	CoursePrice        int    `json:"course_price"`
	CourseQuantity     int    `json:"course_quantity"`
}

type UpdateCourseInput struct {
	CourseCategoryCode string `json:"course_category_code"`
	CourseName         string `json:"course_name"`
	CourseTags         string `json:"course_tags"`
	CoursePrice        int    `json:"course_price"`
	CourseQuantity     int    `json:"course_quantity"`
}
