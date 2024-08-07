package utils

import (
	"github.com/labstack/echo/v4"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SendResponse(c echo.Context, status int, success bool, message string, data interface{}) error {
	return c.JSON(status, Response{
		Status:  status,
		Success: success,
		Message: message,
		Data:    data,
	})
}

func SendError(c echo.Context, status int, message string, data interface{}) error {
	return c.JSON(status, Response{
		Status:  status,
		Success: false,
		Message: message,
		Data:    data,
	})
}

func SendSuccess(c echo.Context, message string, data interface{}) error {
	localize := c.Get("localize").(*i18n.Localizer)
	message = localize.MustLocalize(&i18n.LocalizeConfig{MessageID: message})
	return SendResponse(c, http.StatusOK, true, message, data)
}

func SendCreated(c echo.Context, message string, data interface{}) error {
	localize := c.Get("localize").(*i18n.Localizer)
	message = localize.MustLocalize(&i18n.LocalizeConfig{MessageID: message})
	return SendResponse(c, http.StatusCreated, true, message, data)
}

func SendBadRequest(c echo.Context, message string) error {
	localize := c.Get("localize").(*i18n.Localizer)
	message = localize.MustLocalize(&i18n.LocalizeConfig{MessageID: message})
	return SendError(c, http.StatusBadRequest, message, nil)
}

func SendInternalServerError(c echo.Context, message string) error {
	return SendError(c, http.StatusInternalServerError, message, nil)
}
