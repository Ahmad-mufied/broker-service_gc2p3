package constans

import (
	"github.com/Ahmad-mufied/broker-service_gc2p3/utils"
	"net/http"
)

const (
	ResponseStatusFailed  = "Failed"
	ResponseStatusSuccess = "Success"
)

var (
	ErrNotFound            = utils.NewAPIError(http.StatusNotFound, "Resource not found", nil)
	ErrBadRequest          = utils.NewAPIError(http.StatusBadRequest, "Invalid request data", nil)
	ErrInternalServerError = utils.NewAPIError(http.StatusInternalServerError, "Internal Server Error", nil)
	ErrUnauthorized        = utils.NewAPIError(http.StatusUnauthorized, "Unauthorized access", nil)
	ErrConflict            = utils.NewAPIError(http.StatusConflict, "Resource already exists", nil)
	ErrForbidden           = utils.NewAPIError(http.StatusForbidden, "Forbidden access", nil)
)
