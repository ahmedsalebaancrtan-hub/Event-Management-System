package main

import (
	"fmt"
	"log/slog"

	"github.com/ahmedsaleban/eventManagementsystem/Routes"
	"github.com/ahmedsaleban/eventManagementsystem/infra"
	"github.com/ahmedsaleban/eventManagementsystem/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	slog.Info("initialised enviroment varibale")
	infra.InitEnv()
	config := infra.Configuration
	slog.Info("Connect database successfully")
	infra.DbConnect()
	slog.Info("Connect database succesfully")

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	Routes.RegisterRoute(r)

	slog.Info("application is running successfully on port 5000")
	r.Run(fmt.Sprintf(":%s", config.Port))

}
