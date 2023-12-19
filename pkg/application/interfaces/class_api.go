package interfaces

import "github.com/labstack/echo/v4"

type ClassApi interface {
	Create(e echo.Context) error
}
