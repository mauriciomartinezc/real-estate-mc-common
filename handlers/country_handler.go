package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-common/i18n/locales"
	"github.com/mauriciomartinezc/real-estate-mc-common/services"
	"github.com/mauriciomartinezc/real-estate-mc-common/utils"
)

type CountryHandler struct {
	countryService services.CountryService
}

func NewCountryHandler(countryService services.CountryService) *CountryHandler {
	return &CountryHandler{countryService: countryService}

}

// Countries godoc
// @Summary Get countries
// @Description Get all countries
// @Tags countries
// @Accept json
// @Produce json
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /countries [get]
func (h *CountryHandler) Countries(c echo.Context) error {
	countries, err := h.countryService.Countries()
	if err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	return utils.SendSuccess(c, locales.SuccessResponse, countries)
}
