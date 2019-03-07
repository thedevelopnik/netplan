package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	s "github.com/thedevelopnik/netplan/structs"
)

// CreateNetworkMapEndpoint creates a NetworkMap and returns the created value.
// Returns a 400 if it  can't create the struct,
// or a 500 if the db connection or creation fails.
func (svc netplanService) CreateNetworkMapEndpoint(c *gin.Context) {
	// get the network map object from the request, or send error
	var nm s.NetworkMap
	if err := c.ShouldBindJSON(&nm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
	}

	// create in the db
	if err := svc.repo.CreateNetworkMap(&nm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
	}

	// send full created database object back
	c.JSON(http.StatusCreated, nm)
}

// GetNetworkMapEndpoint gets a NetworkMap struct from a given id.
func (svc netplanService) GetNetworkMapEndpoint(c *gin.Context) {
	id, err := convertParamToInt("id", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
	}

	nm, err := svc.repo.GetNetworkMap(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
	}

	// return it in the response
	c.JSON(http.StatusOK, nm)
}

// UpdateNetworkMapEndpoint updates the name of a
// NetworkMap given an id and name.
func (svc netplanService) UpdateNetworkMapEndpoint(c *gin.Context) {
	// get the values to update with off the request
	var nm s.NetworkMap
	if err := c.ShouldBindJSON(&nm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
	}

	update, err := svc.repo.UpdateNetworkMap(&nm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
	}

	// send back updated value
	c.JSON(http.StatusOK, update)
}

// DeleteNetworkMapEndpoint deletes a NetworkMap given an id.
func (svc netplanService) DeleteNetworkMapEndpoint(c *gin.Context) {
	id, err := convertParamToInt("id", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
	}

	// find db ojbect matching the id
	if err := svc.repo.DeleteNetworkMap(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
	}

	// send back a no content response
	c.Status(http.StatusNoContent)
}
