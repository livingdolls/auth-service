package response

import "github.com/gin-gonic/gin"

// SuccessResponse mengembalikan format sukses
func SuccessResponse(c *gin.Context, data interface{}, message string) {
	c.JSON(200, gin.H{
		"status":  "success",
		"message": message,
		"data":    data,
	})
}

// ErrorResponse mengembalikan format error
func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"status":  "error",
		"message": message,
	})
}
