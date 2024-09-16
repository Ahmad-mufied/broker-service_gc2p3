package handler

import (
	"github.com/Ahmad-mufied/broker-service_gc2p3/config"
	"github.com/Ahmad-mufied/broker-service_gc2p3/constans"
	"github.com/Ahmad-mufied/broker-service_gc2p3/model"
	"github.com/Ahmad-mufied/broker-service_gc2p3/pb"
	"github.com/Ahmad-mufied/broker-service_gc2p3/server/middlewares"
	"github.com/Ahmad-mufied/broker-service_gc2p3/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	UserGrpc pb.AuthUserServiceClient
}

func NewUserHandler(client pb.AuthUserServiceClient) *UserHandler {
	return &UserHandler{
		UserGrpc: client,
	}
}

func (h *UserHandler) Register(c echo.Context) error {
	user := new(model.User)

	err := c.Bind(&user)
	if err != nil {
		return utils.HandleError(c, constans.ErrBadRequest, "Invalid request body")
	}

	// Validate the order struct
	err = config.Validator.Struct(user)
	if err != nil {
		// Format the validation errors
		errors := utils.FormatValidationErrors(err)
		return utils.HandleValidationError(c, errors)
	}

	req := pb.RegisterUserRequest{
		Username: user.Username,
		Password: user.Password,
	}

	res, err := h.UserGrpc.Register(c.Request().Context(), &req)
	if err != nil {
		return utils.HandleError(c, constans.ErrInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, model.JsonResponse{Status: "success", Message: res.Message})
}

func (h *UserHandler) Login(c echo.Context) error {
	user := new(model.User)

	err := c.Bind(&user)
	if err != nil {
		return utils.HandleError(c, constans.ErrBadRequest, "Invalid request body")
	}

	// Validate the order struct
	err = config.Validator.Struct(user)
	if err != nil {
		// Format the validation errors
		errors := utils.FormatValidationErrors(err)
		return utils.HandleValidationError(c, errors)
	}

	// Check if username is exist
	isExistReq := pb.CheckUserIsExistByUsernameRequest{
		Username: user.Username,
	}

	checkUser, err := h.UserGrpc.CheckUserIsExistByUsername(c.Request().Context(), &isExistReq)
	if err != nil {
		return utils.HandleError(c, constans.ErrInternalServerError, err.Error())
	}

	if !checkUser.GetIsExist() {
		return utils.HandleError(c, constans.ErrNotFound, "User Not Found")
	}

	req := pb.LoginUserRequest{
		Username: user.Username,
		Password: user.Password,
	}

	res, err := h.UserGrpc.Login(c.Request().Context(), &req)
	if err != nil {
		return utils.HandleError(c, constans.ErrInternalServerError, err.Error())
	}

	if !res.GetSuccess() {
		return utils.HandleError(c, constans.ErrUnauthorized, res.Message)
	}

	userId := res.GetUserId()
	token, err := middlewares.GenerateToken(userId, user.Username)

	return c.JSON(http.StatusOK, model.JsonResponse{
		Status:  "success",
		Message: res.Message,
		Data:    map[string]interface{}{"token": token},
	})
}

func (h *UserHandler) GetUserById(c echo.Context) error {
	userId := c.Param("id")

	req := pb.GetUserByIdRequest{
		UserId: userId,
	}

	res, err := h.UserGrpc.GetUserById(c.Request().Context(), &req)
	if err != nil {
		return utils.HandleError(c, constans.ErrInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, model.JsonResponse{
		Status:  "success",
		Message: "User found",
		Data:    res.GetUsername(),
	})
}
