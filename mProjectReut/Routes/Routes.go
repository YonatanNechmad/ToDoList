package Routes

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"mProjectReut/Controllers"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AddAllowMethods("OPTIONS")
	corsConfig.AddAllowHeaders("*")

	r.Use(cors.New(corsConfig))
	fmt.Println("Routes :  before creating task group")
	//gin.Context.Header("Content-Type", "application/json")
	task := r.Group("/api/tasks")
	{
		fmt.Println("we are in setupRouter :  Routes")
		task.GET(":id", Controllers.GetATodo)
		task.PATCH(":id", Controllers.UpdateATodo)
		task.DELETE(":id", Controllers.DeleteATodo)
		task.GET(":id/status", Controllers.GetAStatus)
		task.GET(":id/owner", Controllers.GetAOwnerId)
		task.PUT(":id/status", Controllers.ChangeStatus)
		task.PUT(":id/owner", Controllers.ChangeOwner)

	}
	fmt.Println("Routes :  before creating person group")
	person := r.Group("/api/people")
	{
		fmt.Println("we are in setupRouter :  Routes")
		person.POST("", Controllers.CreateAPerson)
		person.GET("", Controllers.GetPersons)
		person.GET(":id", Controllers.GetAPerson)
		person.PATCH(":id", Controllers.UpdateAPerson)
		person.DELETE(":id", Controllers.DeleteAPerson)
		person.POST(":id/tasks", Controllers.AddTask)
		person.GET(":id/tasks", Controllers.GetTasks)

	}

	return r
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT, Delete")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
