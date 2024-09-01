package inputs

type CreateCourseCategoryInput struct {
	CourseCategoryName        string `json:"course_category_name"`
	CourseCategoryDescription string `json:"course_category_description"`
	CourseCategoryTags        string `json:"course_category_tags"`
}

type UpdateCourseCategoryInput struct {
	CourseCategoryName        string `json:"course_category_name"`
	CourseCategoryDescription string `json:"course_category_description"`
	CourseCategoryTags        string `json:"course_category_tags"`
}
