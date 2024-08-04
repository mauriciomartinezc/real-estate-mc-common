package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-common/i18n/locales"
	"github.com/mauriciomartinezc/real-estate-mc-common/service"
	"github.com/mauriciomartinezc/real-estate-mc-common/utils"
)

type CityHandler struct {
	cityService service.CityService
}

func NewCityHandler(e *echo.Group, cityService service.CityService) {
	handler := &CityHandler{cityService: cityService}
	e.GET("/cities/:stateUuid", handler.GetStateCities)
}

func (h *CityHandler) GetStateCities(c echo.Context) error {
	stateUuid := c.Param("stateUuid")
	cities, err := h.cityService.GetStateCities(stateUuid)
	if err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	return utils.SendSuccess(c, locales.SuccessResponse, cities)
}
