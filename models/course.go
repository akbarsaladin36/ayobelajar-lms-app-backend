package models

import "time"

type Course struct {
	CourseId               int    `json:"course_id" gorm:"primaryKey"`
	CourseCategoryCode     string `json:"course_category_code" gorm:"type:varchar(50)"`
	CourseSlugCode         string `json:"course_slug_code" gorm:"type:varchar(100)"`
	CourseName             string `json:"course_name" gorm:"type:varchar(200)"`
	CourseTags             string `json:"course_tags" gorm:"type:text"`
	CoursePrice            int    `json:"course_price" gorm:"type:varchar(100)"`
	CourseQuantity         int    `json:"course_quantity" gorm:"type:varchar(100)"`
	CourseStatusCd         string `json:"course_status_cd" gorm:"type:varchar(50)"`
	CourseCreatedDate      time.Time
	CourserCreatedUserUuid string `json:"course_created_user_uuid" gorm:"type:varchar(200)"`
	CourserCreatedUsername string `json:"course_created_username" gorm:"type:varchar(150)"`
	CourseUpdatedDate      time.Time
	CourseUpdatedUserUuid  string `json:"course_updated_user_uuid" gorm:"type:varchar(200)"`
	CourseUpdatedUsername  string `json:"course_updated_username" gorm:"type:varchar(150)"`
}
