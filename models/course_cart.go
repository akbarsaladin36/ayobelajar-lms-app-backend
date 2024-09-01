package models

import "time"

type CourseCart struct {
	CourseCartId              int    `json:"course_cart_id" gorm:"primaryKey"`
	UserUuid                  string `json:"user_uuid" gorm:"type:varchar(200)"`
	CourseSlugCode            string `json:"course_slug_code" gorm:"type:varchar(100)"`
	CourseCartCode            string `json:"course_cart_code" gorm:"type:varchar(150)"`
	CourseCartQuantity        int    `json:"course_cart_quantity" gorm:"type:varchar(100)"`
	CourseCartCreatedDate     time.Time
	CourseCartCreatedUserUuid string `json:"course_cart_created_user_uuid" gorm:"type:varchar(200)"`
	CourseCartCreatedUsername string `json:"course_cart_created_username" gorm:"type:varchar(150)"`
	CourseCartUpdatedDate     time.Time
	CourseCartUpdatedUserUuid string `json:"course_cart_updated_user_uuid" gorm:"type:varchar(200)"`
	CourseCartUpdatedUsername string `json:"course_cart_updated_username" gorm:"type:varchar(150)"`
}
