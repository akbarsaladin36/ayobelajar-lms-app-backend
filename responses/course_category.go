package responses

import "ayobelajar-app-backend/models"

type CourseCategoriesResponse struct {
	CourseCategoryCode        string `json:"course_category_code"`
	CourseCategoryName        string `json:"course_category_name"`
	CourseCategoryDescription string `json:"course_category_description"`
	CourseCategoryTags        string `json:"course_category_tags"`
}

func ConvertToCourseCategoriesResponse(courseCategoriesRsps models.CourseCategory) CourseCategoriesResponse {
	return CourseCategoriesResponse{
		CourseCategoryCode:        courseCategoriesRsps.CourseCategoryCode,
		CourseCategoryName:        courseCategoriesRsps.CourseCategoryName,
		CourseCategoryDescription: courseCategoriesRsps.CourseCategoryDescription,
		CourseCategoryTags:        courseCategoriesRsps.CourseCategoryTags,
	}
}
