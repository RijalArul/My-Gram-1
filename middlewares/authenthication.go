package middlewares

import (
	"my-gram-1/exceptions"
	"my-gram-1/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenthication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verifyToken, err := helpers.VerifyToken(ctx)

		if err != nil {
			exceptions.Errors(ctx, http.StatusUnauthorized, "Unauthenthicated", "Unauthenthicated")
		}

		ctx.Set("userData", verifyToken)
		ctx.Next()
	}
}
