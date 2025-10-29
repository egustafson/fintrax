package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Credentials struct {
	Username string `json:"name" binding:"required"`
	Password string `json:"pass" binding:"required"`
}

func authHdlr(c *gin.Context) {
	var creds Credentials

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":  creds.Username,
		"token": "xyzzy",
	})
}
