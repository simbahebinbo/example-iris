package routers

import (
	"example-iris/handlers"
	"example-iris/middlewares"
	"github.com/kataras/iris/v12"
)

// Router function handles the iris routers
func Router() {
	router := iris.New()

	router.Get("/login", handlers.Login)
	router.Get("/logout", handlers.Logout)

	userRoute := router.Party("/user", middlewares.CustomAuth)
	userRoute.Get("/get", handlers.GetSessionID)

	router.Run(iris.Addr(":1234"))
}
