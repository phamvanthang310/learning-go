package utils

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func BindAndValidate(e echo.Context, i interface{}) error {
	if err := e.Bind(i); err != nil {
		log.Print(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON request")
	}

	if err := e.Validate(i); err != nil {
		log.Print(err)
		return err
	}

	return nil
}
