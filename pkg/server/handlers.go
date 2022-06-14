package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/ariel17/nirvana-assessment/pkg/services"
)

// CoalesceHandler TODO
// @Summary Shows the status of the application.
// @Description TODO
// @Accept json
// @Produce json
// @Router / [get]
func CoalesceHandler(c *gin.Context) {
	memberIDString := c.Query("member_id")
	if memberIDString == "" {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": "member_id cannot be empty",
		})
		return
	}
	memberID, err := strconv.Atoi(memberIDString)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": "member_id must be a number",
		})
		return
	}
	response, err := services.CoalesceAPIResponses(memberID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]error{"error": err})
		return
	}
	c.JSON(http.StatusOK, response)
}