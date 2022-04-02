package middlewares

import (
	"getcharzp.cn/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthAdminCheck is a middleware function that checks if the user is authenticated with admin role.
func AuthAdminCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		userClaim, err := helper.AnalyseToken(auth)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Unauthorized Authorization",
			})
		}
		if userClaim.IsAdmin != 1 {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Unauthorized Admin",
			})
		}
		c.Next()
	}
}
