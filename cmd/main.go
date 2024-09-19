package main

import (
	"fmt"
	"github.com/Ahmad-mufied/broker-service_gc2p3/config"
	"github.com/Ahmad-mufied/broker-service_gc2p3/gRPC/client"
	"github.com/Ahmad-mufied/broker-service_gc2p3/pb"
	"github.com/Ahmad-mufied/broker-service_gc2p3/server"
	"github.com/Ahmad-mufied/broker-service_gc2p3/server/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
	"log"
)

func main() {
	config.InitViper()
	config.InitValidator()
	userServiceAddress := config.Viper.GetString("USER_SERVICE_ADDRESS")
	bookServiceAddress := config.Viper.GetString("BOOK_SERVICE_ADDRESS")

	userServiceDial, err := grpc.Dial(userServiceAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	bookServiceDial, err := grpc.Dial(bookServiceAddress, grpc.WithInsecure())

	userServicePassword := config.Viper.GetString("USER_SERVICE_PASSWORD")
	refreshDuration := config.Viper.GetDuration("USER_SERVICE_REFRESH_DURATION")

	userServiceAuth := client.NewAuthClient(userServiceDial, config.Viper.GetString("USER_SERVICE_NAME"), userServicePassword)
	userServiceInterceptor, err := client.NewAuthInterceptor(userServiceAuth, client.AuthMethods(), refreshDuration)
	if err != nil {
		log.Fatal(err)
	}

	userServiceDial2, err := grpc.Dial(
		userServiceAddress,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(userServiceInterceptor.Unary()),
		grpc.WithStreamInterceptor(userServiceInterceptor.Stream()),
	)

	if err != nil {
		log.Fatal(err)
	}

	userServiceClient := pb.NewAuthUserServiceClient(userServiceDial2)
	userHandler := handler.NewUserHandler(userServiceClient)

	bookServiceClient := pb.NewBookServiceClient(bookServiceDial)
	bookHandler := handler.NewBookHandler(bookServiceClient)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	handlers := server.NewHandlers(userHandler, bookHandler)
	server.Routes(e, handlers)

	env := config.Viper.GetString("APP_ENV")
	port := "8080"

	if env == "production" {
		log.Println("Running in production mode")
		port = config.Viper.GetString("PORT")
	} else {
		log.Println("Running in development mode")
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%s", port)))

}
