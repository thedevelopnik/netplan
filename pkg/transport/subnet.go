package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	s "github.com/thedevelopnik/netplan/pkg/models"
)

// CreateSubnetEndpoint creates a Subnet and returns the created value.
// Returns a 400 if it  can't create the struct,
// or a 500 if the db connection or creation fails.
func (h netplanHTTP) CreateSubnetEndpoint(c *gin.Context) {
	vpcID, err := convertParamToInt("vpcid", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
	}

	if vpcID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "vpc id parameter must be a positive integer",
		})
	}

	// get the network map object from the request, or send error
	var sn s.Subnet
	if err := c.ShouldBindJSON(&sn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
	}

	sn.VPCID = uint(vpcID)

	// create in the db
	if err := h.svc.CreateSubnet(&sn); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
	}

	// send full created database object back
	c.JSON(http.StatusCreated, sn)
}

// UpdateSubnetEndpoint updates the name of a
// Subnet given an id and name.
func (h netplanHTTP) UpdateSubnetEndpoint(c *gin.Context) {
	// get the values to update with off the request
	var sn s.Subnet
	if err := c.ShouldBindJSON(&sn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
	}

	update, err := h.svc.UpdateSubnet(&sn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
	}

	// send back updated value
	c.JSON(http.StatusOK, update)
}

// DeleteSubnetEndpoint deletes a Subnet given an id.
func (h netplanHTTP) DeleteSubnetEndpoint(c *gin.Context) {
	id, err := convertParamToInt("snid", c)
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
	if err := h.svc.DeleteSubnet(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
	}

	// send back a no content response
	c.Status(http.StatusNoContent)
}
