package handler

import (
	"github.com/Ahmad-mufied/broker-service_gc2p3/config"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RootHandler(c echo.Context) error {

	res := struct {
		Message  string `json:"message"`
		Services []struct {
			Name string `json:"name"`
			Url  string `json:"url"`
			Repo string `json:"repo"`
		} `json:"services"`
	}{
		Message: "Available services",
		Services: []struct {
			Name string `json:"name"`
			Url  string `json:"url"`
			Repo string `json:"repo"`
		}{
			{
				Name: "User Service",
				Url:  config.Viper.GetString("USER_SERVICE_ADDRESS"),
				Repo: config.Viper.GetString("USER_SERVICE_REPO"),
			},
			{
				Name: "Book Service",
				Url:  config.Viper.GetString("BOOK_SERVICE_ADDRESS"),
				Repo: config.Viper.GetString("BOOK_SERVICE_REPO"),
			},
		},
	}

	return c.JSON(http.StatusOK, res)
}
