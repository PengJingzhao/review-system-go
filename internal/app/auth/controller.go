package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRoutes(r *gin.Engine) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", registerUser)
	}

}

func registerUser(c *gin.Context) {
	var registerReq RegisterReq
	if err := c.ShouldBindJSON(&registerReq); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err := RegisterUser(registerReq.Username, registerReq.Password, registerReq.Phone)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
