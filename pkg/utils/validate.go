package utils

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func BindAndValidate(e echo.Context, i interface{}) error {
	if err := e.Bind(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON request")
	}

	if err := e.Validate(i); err != nil {
		return err
	}

	return nil
}
