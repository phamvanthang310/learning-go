package middleware

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strings"
	"student-service/pkg/utils"
)

const AUTH_HEADER = "Authorization"
const CONTEXT_KEY = "claims"

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		headers := c.Request().Header
		token := strings.TrimPrefix(headers.Get(AUTH_HEADER), "Bearer ")

		log.Printf("token: [%s]", token)

		jwtToken, err := utils.ParseAndValidateToken(token)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Missing authorization token")
		}

		if jwtToken.Valid {
			c.Set(CONTEXT_KEY, jwtToken.Claims)
			return next(c)
		}

		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized user")
	}
}
