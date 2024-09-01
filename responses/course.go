package responses

import "ayobelajar-app-backend/models"

type CoursesResponse struct {
	CourseCategoryCode string `json:"course_category_code" gorm:"type:varchar(50)"`
	CourseSlugCode     string `json:"course_slug_code" gorm:"type:varchar(100)"`
	CourseName         string `json:"course_name" gorm:"type:varchar(200)"`
	CourseTags         string `json:"course_tags" gorm:"type:text)"`
	CoursePrice        int    `json:"course_price" gorm:"type:varchar(100)"`
	CourseQuantity     int    `json:"course_quantity" gorm:"type:varchar(100)"`
	CourseStatusCd     string `json:"course_status_cd" gorm:"type:varchar(50)"`
}

type CreateCourseResponse struct {
	CourseName     string `json:"course_name" gorm:"type:varchar(200)"`
	CourseTags     string `json:"course_tags" gorm:"type:text)"`
	CoursePrice    int    `json:"course_price" gorm:"type:varchar(100)"`
	CourseQuantity int    `json:"course_quantity" gorm:"type:varchar(100)"`
	CourseStatusCd string `json:"course_status_cd" gorm:"type:varchar(50)"`
}

type UpdateCourseResponse struct {
	CourseName     string `json:"course_name" gorm:"type:varchar(200)"`
	CourseTags     string `json:"course_tags" gorm:"type:text)"`
	CoursePrice    int    `json:"course_price" gorm:"type:varchar(100)"`
	CourseQuantity int    `json:"course_quantity" gorm:"type:varchar(100)"`
	CourseStatusCd string `json:"course_status_cd" gorm:"type:varchar(50)"`
}

func ConvertToCoursesResponse(coursesRsps models.Course) CoursesResponse {
	return CoursesResponse{
		CourseCategoryCode: coursesRsps.CourseCategoryCode,
		CourseSlugCode:     coursesRsps.CourseSlugCode,
		CourseName:         coursesRsps.CourseName,
		CourseTags:         coursesRsps.CourseTags,
		CoursePrice:        coursesRsps.CoursePrice,
		CourseQuantity:     coursesRsps.CourseQuantity,
		CourseStatusCd:     coursesRsps.CourseStatusCd,
	}
}

func ConvertToCreateCourseResponse(createCourseRsps models.Course) CreateCourseResponse {
	return CreateCourseResponse{
		CourseName:     createCourseRsps.CourseName,
		CourseTags:     createCourseRsps.CourseTags,
		CoursePrice:    createCourseRsps.CoursePrice,
		CourseQuantity: createCourseRsps.CourseQuantity,
		CourseStatusCd: createCourseRsps.CourseStatusCd,
	}
}

func ConvertToUpdateCourseResponse(updateCourseRsps models.Course) UpdateCourseResponse {
	return UpdateCourseResponse{
		CourseName:     updateCourseRsps.CourseName,
		CourseTags:     updateCourseRsps.CourseTags,
		CoursePrice:    updateCourseRsps.CoursePrice,
		CourseQuantity: updateCourseRsps.CourseQuantity,
		CourseStatusCd: updateCourseRsps.CourseStatusCd,
	}
}
