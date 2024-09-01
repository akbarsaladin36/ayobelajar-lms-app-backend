package responses

import "ayobelajar-app-backend/models"

type CourseCartResponse struct {
	UserUuid           string `json:"user_uuid"`
	CourseSlugCode     string `json:"course_slug_code"`
	CourseCartCode     string `json:"course_cart_code"`
	CourseCartQuantity int    `json:"course_cart_quantity"`
}

func ConvertToCourseCartResponse(courseCartRsps models.CourseCart) CourseCartResponse {
	return CourseCartResponse{
		UserUuid:           courseCartRsps.UserUuid,
		CourseSlugCode:     courseCartRsps.CourseSlugCode,
		CourseCartCode:     courseCartRsps.CourseCartCode,
		CourseCartQuantity: courseCartRsps.CourseCartQuantity,
	}
}
