package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-common/i18n/locales"
	"github.com/mauriciomartinezc/real-estate-mc-common/services"
	"github.com/mauriciomartinezc/real-estate-mc-common/utils"
)

type CityHandler struct {
	cityService services.CityService
}

func NewCityHandler(cityService services.CityService) *CityHandler {
	return &CityHandler{cityService: cityService}

}

// GetStateCities godoc
// @Summary Get cities by state
// @Description Get all cities for a specific state
// @Tags cities
// @Accept json
// @Produce json
// @Param stateUuid path string true "State UUID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /cities/{stateUuid} [get]
func (h *CityHandler) GetStateCities(c echo.Context) error {
	stateUuid := c.Param("stateUuid")
	if !utils.IsValidUUID(stateUuid) {
		return utils.SendBadRequest(c, locales.InvalidUuid)
	}
	cities, err := h.cityService.GetStateCities(stateUuid)
	if err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	return utils.SendSuccess(c, locales.SuccessResponse, cities)
}
