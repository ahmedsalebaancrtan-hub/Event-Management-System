package Routes

import (
	"github.com/ahmedsaleban/eventManagementsystem/handlers"
	"github.com/ahmedsaleban/eventManagementsystem/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(r *gin.Engine) {
	ApiGroup := r.Group("/api")

	UserHandler := handlers.RegisterUserHandler()
	UserGroup := ApiGroup.Group("/users")
	{
		UserGroup.POST("/create", UserHandler.CreateUser)
		UserGroup.POST("/login", UserHandler.LoginUser)
		UserGroup.GET("/user/:userId", middleware.Authenticated(), middleware.RequiredRole("ORGANIZER"), UserHandler.GetUserById)
		UserGroup.GET("/allusers", middleware.Authenticated(), middleware.RequiredRole("ADMIN", "STAFF", "ORGANIZER"), UserHandler.GetAllUsers)
	}
}
