package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/theSamuelIyeh/johnethel-school-rms/internals/handlers"
	"github.com/theSamuelIyeh/johnethel-school-rms/internals/middlewares"
)

func InitRouter(e *echo.Echo, handler *handlers.Handler, middleware *middlewares.Middleware) {

	e.GET("/", handler.GetIndexPageHandler)
	e.POST("/theme", handler.SetTheme)

	// dashboard routes
	dashboard := e.Group("/dashboard")
	dashboard.Use(middleware.LoggedinUserMiddleware)
	dashboard.GET("/home", handler.GetDashboardHome)
	dashboard.GET("/sessions", handler.GetDashboardSessions)
	dashboard.GET("/terms", handler.GetDashboardTerms)
	dashboard.GET("/classes", handler.GetDashboardClasses)
	dashboard.GET("/subjects", handler.GetDashboardSubjects)
	dashboard.GET("/students", handler.GetDashboardStudents)
	dashboard.GET("/staff", handler.GetDashboardStaff)
	dashboard.GET("/result-upload", handler.GetDashboardResultUpload)
	dashboard.GET("/psychomotor", handler.GetDashboardPsychomotor)
	dashboard.GET("/comments", handler.GetDashboardComments)
	dashboard.GET("/view-results", handler.GetDashboardViewResults)

	// student routes
	students := e.Group("/students")
	students.Use(middleware.LoggedinUserMiddleware)
	students.GET("/all/:limit/:page", handler.GetAllStudentsHandler)

	// session resource routes
	session := e.Group("/session")
	session.GET("/", handler.GetAllSessionHandler, middleware.LoggedinUserMiddleware)
	session.GET("/:id", handler.GetSessionHandler, middleware.LoggedinUserMiddleware)
	session.POST("/", handler.CreateSessionHandler, middleware.LoggedinUserMiddleware)
	session.PUT("/:id", handler.UpdateSessionHandler, middleware.LoggedinUserMiddleware)
	session.DELETE("/:id", handler.DeleteSessionHandler, middleware.LoggedinUserMiddleware)

	//term resource routes
	term := e.Group("/term")
	// term.GET("/", handler.GetAllTermHandler, middleware.LoggedinUserMiddleware)
	term.GET("/:role", handler.GetTermDropdownHandler, middleware.LoggedinUserMiddleware)
	// term.POST("/", handler.CreateTermHandler, middleware.LoggedinUserMiddleware)
	// term.PUT("/:id", handler.UpdateTermHandler, middleware.LoggedinUserMiddleware)
	// term.DELETE("/:id", handler.DeleteTermHandler, middleware.LoggedinUserMiddleware)

	// term class resource routes
	termClass := e.Group("/term-class")
	termClass.GET("/", handler.GetAllTermClassHandler, middleware.LoggedinUserMiddleware)
	termClass.GET("/:id", handler.GetTermClassHandler, middleware.LoggedinUserMiddleware)
	termClass.POST("/", handler.CreateTermClassHandler, middleware.LoggedinUserMiddleware)
	termClass.PUT("/:id", handler.UpdateTermClassHandler, middleware.LoggedinUserMiddleware)
	termClass.DELETE("/:id", handler.DeleteTermClassHandler, middleware.LoggedinUserMiddleware)

	// term subject resource routes
	termSubject := e.Group("/term-subject")
	termSubject.GET("/", handler.GetAllTermSubjectHandler, middleware.LoggedinUserMiddleware)
	termSubject.GET("/:id", handler.GetTermSubjectHandler, middleware.LoggedinUserMiddleware)
	termSubject.POST("/", handler.CreateTermSubjectHandler, middleware.LoggedinUserMiddleware)
	termSubject.PUT("/:id", handler.UpdateTermSubjectHandler, middleware.LoggedinUserMiddleware)
	termSubject.DELETE("/:id", handler.DeleteTermSubjectHandler, middleware.LoggedinUserMiddleware)

	// auth routes
	auth := e.Group("/auth")
	auth.GET("/login", handler.GetLoginFormHandler, middleware.LoggedOutUserMiddleware)
	auth.POST("/login", handler.LoginHandler, middleware.LoggedOutUserMiddleware)
	// auth.GET("/updatepassword", handler.GetUpdatePasswordFormHandler, middleware.LoggedinUserMiddleware)
	auth.PUT("/updatepassword", handler.UpdatePasswordHandler, middleware.LoggedinUserMiddleware)
	auth.GET("/forgotpassword", handler.GetForgotPasswordFormHandler, middleware.LoggedOutUserMiddleware)
	auth.PUT("/forgotpassword", handler.GetForgotPasswordFormHandler, middleware.LoggedOutUserMiddleware)
	auth.POST("/logout", handler.LogoutHandler, middleware.LoggedinUserMiddleware)
}
