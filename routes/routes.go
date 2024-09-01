package routes

import (
	"ayobelajar-app-backend/config"
	"ayobelajar-app-backend/handlers"
	"ayobelajar-app-backend/middleware"
	"ayobelajar-app-backend/repositories"
	"ayobelajar-app-backend/services"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ConnectRoutes() {

	authRepositories := repositories.NewAuthRepository(config.DB)
	authServices := services.NewAuthService(authRepositories)
	authHandler := handlers.NewAuthHandler(authServices)

	userRepositories := repositories.NewUserRepository(config.DB)
	userServices := services.NewUserService(userRepositories)
	userHandler := handlers.NewUserHandler(userServices)

	courseRepositories := repositories.NewCourseRepository(config.DB)
	courseServices := services.NewCourseService(courseRepositories)
	courseHandler := handlers.NewCourseHandler(courseServices)

	courseCategoryRepositories := repositories.NewCourseCategoryRepository(config.DB)
	courseCategoryServices := services.NewCourseCategoryService(courseCategoryRepositories)
	courseCategoryHandler := handlers.NewCourseCategoryHandler(courseCategoryServices)

	courseCartRepositories := repositories.NewCourseCartRepository(config.DB)
	courseCartServices := services.NewCourseCartService(courseCartRepositories)
	courseCartHandler := handlers.NewCourseCartHandler(courseCartServices)

	courseInvoicePaymentRepositories := repositories.NewCourseInvoicePaymentRepository(config.DB)
	courseInvoicePaymentServices := services.NewCourseInvoicePaymentService(courseInvoicePaymentRepositories)
	courseInvoicePaymentHandler := handlers.NewCourseInvoicePaymentHandler(courseInvoicePaymentServices)

	router := gin.Default()

	router.Use(cors.Default())

	v1 := router.Group("/ayo-belajar-lms-api/v1")

	v1.POST("/auth/register", authHandler.RegisterUserHandler)
	v1.POST("/auth/login", authHandler.LoginUserHandler)

	v1Admin := router.Group("/ayo-belajar-lms-api/v1/admin").Use(middleware.JWTAuthMiddleware(), middleware.IsAdminAccess())

	v1Admin.GET("/users", userHandler.GetAllUsersHandler)
	v1Admin.GET("/users/:username", userHandler.GetOneUserHandler)
	v1Admin.POST("/users", userHandler.CreateNewUserHandler)
	v1Admin.PATCH("/users/:username", userHandler.UpdateOneUserHandler)
	v1Admin.DELETE("/users/:username", userHandler.DeleteOneUserHandler)

	v1Admin.GET("/courses", courseHandler.GetAllCoursesHandler)
	v1Admin.GET("/courses/:course_slug_code", courseHandler.GetOneCourseHandler)
	v1Admin.POST("/courses", courseHandler.CreateNewCourseHandler)
	v1Admin.PATCH("/courses/:course_slug_code", courseHandler.UpdateOneCourseHandler)
	v1Admin.DELETE("/courses/:course_slug_code", courseHandler.DeleteOneCourseHandler)

	v1Admin.GET("/course-categories", courseCategoryHandler.GetAllCourseCategoriesHandler)
	v1Admin.GET("/course-categories/:course_category_code", courseCategoryHandler.GetOneCourseCategoryHandler)
	v1Admin.POST("/course-categories", courseCategoryHandler.CreateNewCourseCategoryHandler)
	v1Admin.PATCH("/course-categories/:course_category_code", courseCategoryHandler.UpdateOneCourseCategoryHandler)
	v1Admin.DELETE("/course-categories/:course_category_code", courseCategoryHandler.DeleteOneCourseCategoryHandler)

	v1Admin.GET("/course-cart", courseCartHandler.GetAllCourseCartsHandler)
	v1Admin.GET("/course-cart/:course_cart_code", courseCartHandler.GetCourseCartHandler)

	v1Admin.GET("/course-invoice-payments", courseInvoicePaymentHandler.GetAllCourseInvoicePaymentsHandler)
	v1Admin.GET("/course-invoice-payments/:course_invoice_payment_code", courseInvoicePaymentHandler.GetOneCourseInvoicePaymentHandler)

	v1User := router.Group("/ayo-belajar-lms-api/v1/user").Use(middleware.JWTAuthMiddleware(), middleware.IsUserAccess())

	v1User.GET("/course-cart/my-course-carts", courseCartHandler.GetMyCourseCartsHandler)
	v1User.GET("/course-cart/:course_cart_code", courseCartHandler.GetCourseCartHandler)
	v1User.POST("/course-cart", courseCartHandler.CreateNewCourseCartHandler)
	v1User.PATCH("/course-cart/:course_cart_code", courseCartHandler.UpdateOneCourseCartHandler)
	v1User.DELETE("/course-cart/:course_cart_code", courseCartHandler.DeleteOneCourseCartHandler)

	v1User.GET("/course-invoice-payments/my-course-invoice-payments", courseInvoicePaymentHandler.GetMyCourseInvoicePaymentsHandler)
	v1User.POST("/course-invoice-payments", courseInvoicePaymentHandler.CreateNewCourseInvoicePaymentHandler)

	router.Run(os.Getenv("APP_PORT"))
}
