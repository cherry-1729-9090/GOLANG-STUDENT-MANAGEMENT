package routes

import (
    "github.com/gin-gonic/gin"
    "lms-backend/controllers"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

    api := router.Group("/api")
    {
        api.POST("/register", controllers.Register)
        api.POST("/login", controllers.Login)

        api.GET("/courses", controllers.GetCourses)
        api.POST("/courses", controllers.CreateCourse)

        api.GET("/batches", controllers.GetBatches)
        api.POST("/batches", controllers.CreateBatch)

        api.GET("/classes", controllers.GetClasses)
        api.POST("/classes", controllers.CreateClass)

        api.GET("/assignments", controllers.GetAssignments)
        api.POST("/assignments", controllers.CreateAssignment)

        api.GET("/questions", controllers.GetQuestions)
        api.POST("/questions", controllers.CreateQuestion)

        api.GET("/students", controllers.GetStudents)
        api.POST("/students", controllers.CreateStudent)

        api.GET("/teachers", controllers.GetTeachers)
        api.POST("/teachers", controllers.CreateTeacher)
    }
    return router
}
