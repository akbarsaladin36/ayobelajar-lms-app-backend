package inputs

type CreateCourseCartInput struct {
	CourseSlugCode     string `json:"course_slug_code"`
	CourseCartCode     string `json:"course_cart_code"`
	CourseCartQuantity int    `json:"course_cart_quantity"`
}

type UpdateCourseCartInput struct {
	CourseSlugCode     string `json:"course_slug_code"`
	CourseCartCode     string `json:"course_cart_code"`
	CourseCartQuantity int    `json:"course_cart_quantity"`
}
