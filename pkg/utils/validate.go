package utils

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func BindAndValidate(i interface{}, e echo.Context) error {
	if err := e.Bind(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON request")
	}

	if err := e.Validate(i); err != nil {
		return err
	}

	return nil
}
