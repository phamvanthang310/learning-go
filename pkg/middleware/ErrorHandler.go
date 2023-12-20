package middleware

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func ErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	message := "Internal Server Error"

	var httpErr *echo.HTTPError
	if errors.As(err, &httpErr) {
		code = httpErr.Code
		message = fmt.Sprintf("%v", httpErr.Message)
		log.Printf("Error: %s", message)
	}

	response := ErrorResponse{
		message,
		code,
	}

	c.JSON(code, &response)
}
