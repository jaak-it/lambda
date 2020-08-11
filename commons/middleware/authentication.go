package middleware

import (
	"github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
	"github.com/jaak-it/lambda/commons/models"
    "github.com/sirupsen/logrus"
    "net/http"
    "strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.GetHeader("Authorization")
		if clientToken == "" {
			logrus.Errorf("Authorization token was not provided")
			c.JSON(
				http.StatusUnauthorized,
				models.ResponseError{
					Response: models.Response{
						Message: "Authorization Token is required",
					},
				},
			)
			c.Abort()
			return
		}

		extractedToken := strings.Split(clientToken, "Bearer ")

		if len(extractedToken) == 2 {
			clientToken = strings.TrimSpace(extractedToken[1])
		} else {
			logrus.Errorf("Incorrect Format of Auth Token")
			c.JSON(
				http.StatusBadRequest,
                models.ResponseError{
					Response: models.Response{
						Message: "Incorrect Format of Authorization Token",
					},
				},
			)
			c.Abort()
			return
		}

		// TODO verificar el error de la firma del token
		parsedToken, _ := jwt.ParseWithClaims(clientToken, &models.ClaimsToken{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(""), nil
		})
		claims, _ := parsedToken.Claims.(*models.ClaimsToken)

		if claims.Username == "" {
			c.JSON(
				http.StatusBadRequest,
                models.ResponseError{
					Response: models.Response{
						Message: "Token claims is invalid",
					},
				},
			)
			c.Abort()
			return
		}


		c.Set("claims", claims)
		c.Next()
	}
}
