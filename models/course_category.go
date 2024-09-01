package models

import "time"

type CourseCategory struct {
	CourseCategoryId              int    `json:"course_category_id" gorm:"primaryKey"`
	CourseCategoryCode            string `json:"course_category_code" gorm:"type:varchar(100)"`
	CourseCategoryName            string `json:"course_category_name" gorm:"type:varchar(200)"`
	CourseCategoryDescription     string `json:"course_category_description" gorm:"type:text"`
	CourseCategoryTags            string `json:"course_category_tags" gorm:"type:text"`
	CourseCategoryStatusCd        string `json:"course_category_status_cd" gorm:"type:varchar(50)"`
	CourseCategoryCreatedDate     time.Time
	CourseCategoryCreatedUserUuid string `json:"course_category_created_user_uuid" gorm:"type:varchar(200)"`
	CourseCategoryCreatedUsername string `json:"course_category_created_username" gorm:"type:varchar(150)"`
	CourseCategoryUpdatedDate     time.Time
	CourseCategoryUpdatedUserUuid string `json:"course_category_updated_user_uuid" gorm:"type:varchar(200)"`
	CourseCategoryUpdatedUsername string `json:"course_category_updated_username" gorm:"type:varchar(150)"`
}
