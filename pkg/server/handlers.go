package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/ariel17/nirvana-assessment/pkg/services"
)

// CoalesceHandler takes a member_id parameter to pass thru external APIs. It
// validates the correctness of input and response values, returning correct
// status codes.
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