package utils

import (
	"errors"
	"github.com/go-playground/validator/v10"
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
	return SendResponse(c, http.StatusBadRequest, false, message, nil)
}

func SendErrorValidations(c echo.Context, message string, err error) error {
	localize := c.Get("localize").(*i18n.Localizer)
	message = localize.MustLocalize(&i18n.LocalizeConfig{MessageID: message})
	var validationErrors validator.ValidationErrors
	errors.As(err, &validationErrors)
	return SendResponse(c, http.StatusUnprocessableEntity, false, message, FormatValidationErrors(localize, validationErrors))
}

func SendInternalServerError(c echo.Context, message string) error {
	return SendResponse(c, http.StatusInternalServerError, false, message, nil)
}
