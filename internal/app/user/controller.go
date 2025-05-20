package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func RegisterRoutes(r *gin.Engine) {
	userGroup := r.Group("/users")
	{
		userGroup.GET("/listUsers", listUsers)
		userGroup.GET("/getUser/:id", getUser)
		userGroup.POST("/createUser", createUser)
	}
}

func listUsers(c *gin.Context) {
	users, err := ListUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func getUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	user, err := GetUser(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	err := CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"userId": user.ID})
}
