package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/fsyabc/test/internal/service"
)

// ScanHandler handles the port scan requests.
func ScanHandler(c *gin.Context) {
	target := c.Query("target")
	if target == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Target is required"})
		return
	}

	result, err := service.ScanPorts(target)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}

