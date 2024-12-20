package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-common/i18n/locales"
	"github.com/mauriciomartinezc/real-estate-mc-common/services"
	"github.com/mauriciomartinezc/real-estate-mc-common/utils"
)

type StateHandler struct {
	stateService services.StateService
}

func NewStateHandler(stateService services.StateService) *StateHandler {
	return &StateHandler{stateService: stateService}
}

// GetCountryStates godoc
// @Summary Get states by country
// @Description Get all states for a specific country
// @Tags states
// @Accept json
// @Produce json
// @Param countryUuid path string true "Country UUID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /states/{countryUuid} [get]
func (h *StateHandler) GetCountryStates(c echo.Context) error {
	countryUuid := c.Param("countryUuid")
	if !utils.IsValidUUID(countryUuid) {
		return utils.SendBadRequest(c, locales.InvalidUuid)
	}
	states, err := h.stateService.GetCountryStates(countryUuid)
	if err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	return utils.SendSuccess(c, locales.SuccessResponse, states)
}
