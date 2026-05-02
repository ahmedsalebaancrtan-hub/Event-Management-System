package Routes

import (
	"github.com/ahmedsaleban/eventManagementsystem/handlers"
	"github.com/ahmedsaleban/eventManagementsystem/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(r *gin.Engine) {
	ApiGroup := r.Group("/api")

	UserHandler := handlers.RegisterUserHandler()
	EventHandler := handlers.RegisterEventHandler()
	UserGroup := ApiGroup.Group("/users")

	{
		UserGroup.POST("/create", UserHandler.CreateUser)
		UserGroup.POST("/login", UserHandler.LoginUser)
		UserGroup.POST("/verify-2fa-login", UserHandler.Verify2FALogin)
		UserGroup.GET("/user/:userId", middleware.Authenticated(), middleware.RequiredRole("ORGANIZER"), UserHandler.GetUserById)
		UserGroup.GET("/allusers", middleware.Authenticated(), middleware.RequiredRole("ADMIN", "STAFF", "ORGANIZER"), UserHandler.GetAllUsers)
		UserGroup.GET("/whoami", middleware.Authenticated(), middleware.RequiredRole("ADMIN", "ORGANIZER", "STAFF"), UserHandler.WhoAmI)
		UserGroup.POST("/Refresh_token", middleware.RefreshAuthenticated(), UserHandler.RefreshToken)
		UserGroup.POST("/forget-password", UserHandler.ForgotPassword)
		UserGroup.POST("/reset-password", UserHandler.ResetPassword)
		UserGroup.POST("/admin/reset-password", middleware.Authenticated(), middleware.RequiredRole("ADMIN"), UserHandler.ResetPasswordByAdmin)
	}

	EventGroup := ApiGroup.Group("/events")
	{
		EventGroup.POST("/create", middleware.Authenticated(), middleware.RequiredRole("ADMIN", "STAFF"), EventHandler.CreateEvent)
		EventGroup.GET("/list", middleware.Authenticated(), middleware.RequiredRole("ADMIN", "STAFF", "ORGANIZER"), EventHandler.Getall)
		EventGroup.GET("/details/:event_id", middleware.Authenticated(), middleware.RequiredRole("ADMIN", "STAFF", "ORGANIZER"), EventHandler.FindEventByid)
		EventGroup.PATCH("/Update/:id", middleware.Authenticated(), middleware.RequiredRole("ADMIN", "STAFF", "ORGANIZER"), EventHandler.UpdateEvent)

	}

}
