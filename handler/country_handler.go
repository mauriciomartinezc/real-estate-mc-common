package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-common/i18n/locales"
	"github.com/mauriciomartinezc/real-estate-mc-common/service"
	"github.com/mauriciomartinezc/real-estate-mc-common/utils"
)

type CountryHandler struct {
	countryService service.CountryService
}

func NewCountryHandler(e *echo.Group, countryService service.CountryService) {
	handler := &CountryHandler{countryService: countryService}
	e.GET("/countries", handler.Countries)
}

func (h *CountryHandler) Countries(c echo.Context) error {
	countries, err := h.countryService.Countries()
	if err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	return utils.SendSuccess(c, locales.SuccessResponse, countries)
}
