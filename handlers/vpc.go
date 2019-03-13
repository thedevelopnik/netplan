package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	s "github.com/thedevelopnik/netplan/structs"
)

// CreateVPCEndpoint creates a VPC and returns the created value.
// Returns a 400 if it  can't create the struct,
// or a 500 if the db connection or creation fails.
func (svc netplanService) CreateVPCEndpoint(c *gin.Context) {
	nmID, err := convertParamToInt("nmid", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
	}

	// get the vpc object from the request, or send error
	var vpc s.VPC
	if err := c.ShouldBindJSON(&vpc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
	}
	vpc.NetworkMapID = nmID

	// create in the db
	if err := svc.repo.CreateVPC(&vpc); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
	}

	// send full created database object back
	c.JSON(http.StatusCreated, vpc)
}

// UpdateVPCEndpoint updates the name of a
// VPC given an id and name.
func (svc netplanService) UpdateVPCEndpoint(c *gin.Context) {
	// get the values to update with off the request
	var vpc s.VPC
	if err := c.ShouldBindJSON(&vpc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
	}

	// save in the db or send error
	update, err := svc.repo.UpdateVPC(&vpc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
	}

	// send back updated value
	c.JSON(http.StatusOK, update)
}

// DeleteVPCEndpoint deletes a VPC given an id.
func (svc netplanService) DeleteVPCEndpoint(c *gin.Context) {
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
	if err := svc.repo.DeleteVPC(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
	}

	// send back a no content response
	c.Status(http.StatusNoContent)
}
