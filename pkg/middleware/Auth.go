package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strings"
	"student-service/pkg/application/model"
	"student-service/pkg/constant"
	"student-service/pkg/utils"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		headers := c.Request().Header
		token := strings.TrimPrefix(headers.Get(constant.AuthHeader), "Bearer ")

		log.Printf("token: [%s]", token)

		jwtToken, err := utils.ParseAndValidateToken(token)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Missing authorization token")
		}

		if jwtToken.Valid {
			rawClaims := jwtToken.Claims.(jwt.MapClaims)
			c.Set(constant.Claims, &model.AuthClaims{
				ID:       rawClaims["id"].(string),
				Name:     rawClaims["name"].(string),
				Username: rawClaims["userName"].(string),
			})
			return next(c)
		}

		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized user")
	}
}
