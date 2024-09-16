package main

import (
	"github.com/Ahmad-mufied/broker-service_gc2p3/config"
	"github.com/Ahmad-mufied/broker-service_gc2p3/gRPC/clientAuth"
	"github.com/Ahmad-mufied/broker-service_gc2p3/pb"
	"github.com/Ahmad-mufied/broker-service_gc2p3/server/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
	"log"
)

func main() {
	config.InitViper()
	config.InitValidator()

	serverAddress := "localhost:50051"
	log.Println("Server running on port :50001")

	cc1, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	userServicePassword := config.Viper.GetString("USER_SERVICE_PASSWORD")
	refreshDuration := config.Viper.GetDuration("REFRESH_DURATION")

	userServiceAuth := clientAuth.NewAuthClient(cc1, config.Viper.GetString("USER_SERVICE_NAME"), userServicePassword)
	userServiceInterceptor, err := clientAuth.NewAuthInterceptor(userServiceAuth, clientAuth.AuthMethods(), refreshDuration)
	if err != nil {
		log.Fatal(err)
	}

	cc2, err := grpc.Dial(
		serverAddress,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(userServiceInterceptor.Unary()),
		grpc.WithStreamInterceptor(userServiceInterceptor.Stream()),
	)

	if err != nil {
		log.Fatal(err)
	}

	userServiceClient := pb.NewAuthUserServiceClient(cc2)
	userHandler := handler.NewUserHandler(userServiceClient)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/register", userHandler.Register)
	err = e.Start(":8080")
	if err != nil {
		e.Logger.Fatal(err)
	}

}
