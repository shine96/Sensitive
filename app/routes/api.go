package routes

import (
	"Sensitive/app/controllers"
	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	r := gin.Default()
	r.GET("/", controllers.Index)
	api := r.Group("/api")
	{
		api.POST("/check", controllers.Check)
		api.POST("/wordAdd", controllers.Add)
		api.POST("/save", controllers.Save)
		api.GET("/list", controllers.List)
	}
	project := r.Group("/api/project")
	{
		project.POST("/create", controllers.CreateProject)
		project.POST("/check", controllers.ProjectCheck)
		project.POST("/del", controllers.DelProject)
		project.POST("/addWord", controllers.ProjectAddWord)
		project.POST("/saveAll", controllers.SaveAll)
		project.GET("/list", controllers.DictList)
	}

	return r
}
