package server

import (
	"github.com/Ahmad-mufied/broker-service_gc2p3/config"
	"github.com/Ahmad-mufied/broker-service_gc2p3/server/handler"
	"github.com/Ahmad-mufied/broker-service_gc2p3/server/middlewares"
	"github.com/golang-jwt/jwt/v5"
	echoJWT "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Handlers struct {
	user *handler.UserHandler
	book *handler.BookHandler
}

func NewHandlers(userHandler *handler.UserHandler, bookHandler *handler.BookHandler) *Handlers {
	return &Handlers{
		user: userHandler,
		book: bookHandler,
	}
}

func Routes(e *echo.Echo, handlers *Handlers) {

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", handler.RootHandler)

	//Register Swagger route
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/users/register", handlers.user.Register)
	e.POST("/users/login", handlers.user.Login)

	jwtConfig := echoJWT.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(middlewares.JWTCustomClaims)
		},
		SigningKey: []byte(config.Viper.GetString("JWT_SECRET")),
	}

	// Book Routes
	e.POST("/books", handlers.book.CreateBook, echoJWT.WithConfig(jwtConfig))
	e.GET("/book/:id", handlers.book.GetBookById, echoJWT.WithConfig(jwtConfig))
	e.PUT("/book/:id", handlers.book.UpdateBook, echoJWT.WithConfig(jwtConfig))
	e.DELETE("/book/:id", handlers.book.DeleteBook, echoJWT.WithConfig(jwtConfig))

}
