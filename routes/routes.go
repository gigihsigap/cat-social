package routes

import (
	"github.com/gin-gonic/gin"
	"cat-social/controllers"
	"cat-social/middlewares/auth"
	"cat-social/repositories"
	"cat-social/services"
	"github.com/jackc/pgx/v5"
)

func SetupRouter(conn *pgx.Conn) *gin.Engine {

	// Repository setup

	userRepository := repository.NewUserRepository(conn)
	catRepository := repository.NewCatRepository(conn)
	matchRepository := repository.NewMatchRepository(conn)

	// Controller & Service setup

	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	matchService := service.NewMatchService(matchRepository, catRepository)
	matchController := controller.NewMatchController(matchService)

	catService := service.NewCatService(catRepository, matchService)
	catController := controller.NewCatController(catService)

	// Router setup

	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	routerV1 := router.Group("/v1")

	userV1 := routerV1.Group("/user")
	userV1.POST("/register", userController.SignUp)
	userV1.POST("/login", userController.SignIn)

	catRouter := routerV1.Group("/cat", auth.RequireAuth)
	catRouter.GET("/", catController.FindAll)
	catRouter.POST("/", catController.Create)
	catRouter.PUT("/:id", catController.Update)
	catRouter.GET("/:id", catController.FindByID)
	catRouter.DELETE("/:id", catController.Delete)

	matchRouter := routerV1.Group("/match", auth.RequireAuth)
	matchRouter.POST("/", matchController.Create)
	matchRouter.POST("/approve", matchController.Approve)
	matchRouter.POST("/reject", matchController.Reject)


	return router
}
