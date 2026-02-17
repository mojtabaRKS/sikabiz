package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uh *UserHandler) GetUser(c *gin.Context) {
	userId := c.Param("id")

	res, err := uh.userService.GetUser(c, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    res,
	})
}
