package Routes

import (
	"github.com/ahmedsaleban/eventManagementsystem/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(r *gin.Engine) {
	ApiGroup := r.Group("/api")

	UserHandler := handlers.RegisterUserHandler()
	UserGroup := ApiGroup.Group("/users")
	{
		UserGroup.POST("/create", UserHandler.CreateUser)
		UserGroup.POST("/login", UserHandler.LoginUser)
	}
}
