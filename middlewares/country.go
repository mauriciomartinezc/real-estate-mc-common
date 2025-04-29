package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-common/i18n/locales"
	"github.com/mauriciomartinezc/real-estate-mc-common/utils"
)

func CompanyHandler() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			companyId := c.Request().Header.Get("X-Company-Id")
			if companyId == "" {
				return utils.SendBadRequest(c, locales.MissingCompanyHeader)
			}
			c.Set("companyId", companyId)
			return next(c)
		}
	}
}
