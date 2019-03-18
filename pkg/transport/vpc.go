package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	s "github.com/thedevelopnik/netplan/pkg/models"
)

// CreateVPCEndpoint creates a VPC and returns the created value.
// Returns a 400 if it  can't create the struct,
// or a 500 if the db connection or creation fails.
func (h netplanHTTP) CreateVPCEndpoint(c *gin.Context) {
	nmID, err := convertParamToInt("nmid", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
	}

	if nmID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id parameter must be a positive integer",
		})
	}

	// get the vpc object from the request, or send error
	var vpc s.VPC
	if err := c.ShouldBindJSON(&vpc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
	}
	vpc.NetworkMapID = uint(nmID)

	// create in the db
	if err := h.svc.CreateVPC(&vpc); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
	}

	// send full created database object back
	c.JSON(http.StatusCreated, vpc)
}

// UpdateVPCEndpoint updates the name of a
// VPC given an id and name.
func (h netplanHTTP) UpdateVPCEndpoint(c *gin.Context) {
	// get the values to update with off the request
	var vpc s.VPC
	if err := c.ShouldBindJSON(&vpc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
	}

	// save in the db or send error
	update, err := h.svc.UpdateVPC(&vpc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
	}

	// send back updated value
	c.JSON(http.StatusOK, update)
}

// DeleteVPCEndpoint deletes a VPC given an id.
func (h netplanHTTP) DeleteVPCEndpoint(c *gin.Context) {
	id, err := convertParamToInt("id", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
	}

	if id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id parameter must be a positive integer",
		})
	}

	// delete the object
	if err := h.svc.DeleteVPC(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
	}

	// send back a no content response
	c.Status(http.StatusNoContent)
}
