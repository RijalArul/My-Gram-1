package accepted_responses

import "github.com/gin-gonic/gin"

func SuccessResponse(ctx *gin.Context, code int, message string, data interface{}) {
	switch code != 0 {
	case code == 201:
		ctx.JSON(code, gin.H{
			"message": message,
			"data":    data,
		})
	case code == 200 && message == "Login Success":
		ctx.JSON(code, gin.H{
			"message": message,
			"token":   data,
		})
	case code == 200:
		ctx.JSON(code, gin.H{
			"message": message,
			"data":    data,
		})
	default:
		ctx.JSON(code, gin.H{
			"message": message,
		})
	}
}
