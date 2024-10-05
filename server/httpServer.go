package server

import (
	"database/sql"
	"log"
	"mutualfund/controllers"
	"mutualfund/repositories"
	"mutualfund/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config            *viper.Viper
	router            *gin.Engine
	usersController *controllers.UsersController
	fundSchemeController *controllers.FundSchemeController
	orderController *controllers.OrderController
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
	usersRepository := repositories.NewUsersRepository(dbHandler)
	usersService := services.NewUsersService(usersRepository)
	usersController := controllers.NewUsersController(usersService)

	orderRepository := repositories.NewOrderRepository(dbHandler)
	orderService := services.NewOrderService(orderRepository)
	orderController := controllers.NewOrderController(orderService)

	fundSchemeRepository := repositories.NewFundSchemeRepository(dbHandler)
	fundSchemeService := services.NewFundSchemeService(fundSchemeRepository,orderRepository)
	fundSchemeController := controllers.NewFundSchemeController(fundSchemeService)


	router := gin.Default()

	router.POST("/user", usersController.CreateUser)
	router.POST("/user/login", usersController.LoginUser)

	router.POST("/fundscheme",fundSchemeController.CreateFundScheme)
	router.GET("/fundscheme", fundSchemeController.GetFundSchemes)
	router.GET("/fundscheme/:id", fundSchemeController.GetFundSchemeByID)
	router.PUT("/fundscheme/:id", fundSchemeController.UpdateFundScheme)
	router.DELETE("/fundscheme/:id", fundSchemeController.DeleteFundScheme)
	router.GET("/fundscheme/filter", fundSchemeController.GetFilteredFundSchemes)

	router.POST("/user/placeorder",orderController.CreateOrder)
	router.GET("/user/:id/orders",orderController.GetOrdersByUserID)
	
	return HttpServer{
		config:            config,
		router:            router,
		usersController: usersController,
		fundSchemeController: fundSchemeController,
		orderController: orderController,
	}
}

func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatalf("Error while starting HTTP server: %v", err)
	}
}
