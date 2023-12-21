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

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
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
				Role:     rawClaims["role"].(string),
			})
			return next(c)
		}

		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized user")
	}
}

func Authorization(roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims, err := utils.GetTokenClaims(c)

			if err != nil {
				return err
			}

			for _, v := range roles {
				if v == claims.Role {
					return next(c)
				}
			}

			return echo.ErrForbidden
		}
	}
}
