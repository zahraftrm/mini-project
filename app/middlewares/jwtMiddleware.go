package middlewares

import (
	"net/http"
	"time"

	"github.com/zahraftrm/mini-project/app/config"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(config.SECRET_JWT),
		SigningMethod: "HS256",
	})
}

func CreateToken(Id int, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true

	if role == "admin" {
		claims["adminId"] = Id
	} else if role == "teacher" {
		claims["teacherId"] = Id
	}

	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.SECRET_JWT))
}

func ExtractTokenInfo(e echo.Context, roleToExtract ...string) (int, string) {
	// Mengekstrak token berdasarkan peran yang diberikan
	var token *jwt.Token
	for _, role := range roleToExtract {
		token = e.Get(role).(*jwt.Token)
		if token.Valid {
			break
		}
	}

	if token == nil || !token.Valid {
		return 0, ""
	}

	claims := token.Claims.(jwt.MapClaims)
	role := claims["role"].(string)
	var userId int

	if role == "admin" {
		userId = int(claims["adminId"].(float64))
	} else if role == "teacher" {
		userId = int(claims["teacherId"].(float64))
	} else {
		return 0, ""
	}

	return userId, role
}


func RoleAuthorization(roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			_, extractedRole := ExtractTokenInfo(c, roles...)

			// Memeriksa role
			authorized := false
			for _, allowedRole := range roles {
				if extractedRole == allowedRole {
					authorized = true
					break
				}
			}

			if !authorized {
				return c.JSON(http.StatusForbidden, "Permission denied")
			}

			return next(c)
		}
	}
}

