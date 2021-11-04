package middlewares

import (
	"hermes/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			utils.HandleErrorResponse(c, http.StatusBadRequest, c.Request.Method, "no token provided")
			return
		}
		token := header[len("Bearer "):]
		validated, err := utils.NewJWTService().Validate(token)
		if err != nil {
			utils.HandleErrorResponse(c, http.StatusInternalServerError, c.Request.Method, "token not authorized")
			return
		}
		claim, ok := validated.Claims.(jwt.MapClaims)
		if !ok && !validated.Valid {
			utils.HandleErrorResponse(c, http.StatusInternalServerError, c.Request.Method, "token not authorized")
			return
		}
		user := claim["user_id"]
		c.Set("user", user)
		c.Next()
	}
}
