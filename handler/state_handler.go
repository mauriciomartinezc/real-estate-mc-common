package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-common/i18n/locales"
	"github.com/mauriciomartinezc/real-estate-mc-common/service"
	"github.com/mauriciomartinezc/real-estate-mc-common/utils"
)

type StateHandler struct {
	stateService service.StateService
}

func NewStateHandler(e *echo.Group, stateService service.StateService) {
	handler := &StateHandler{stateService: stateService}
	e.GET("/states/:countryUuid", handler.GetCountryStates)
}

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