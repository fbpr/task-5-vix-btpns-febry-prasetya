package middlewares
import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/fbpr/task-5-vix-btpns-febry-prasetya/helpers"
)
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader("Authorization")
		if accessToken == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "need access token"})
			return
		}

		verifyToken, err := helpers.ValidateJWT(strings.Split(accessToken, "Bearer ")[1]) 
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}

		c.Set("userData", verifyToken)
		c.Next()
	}
}