package transport

import (
	"net/http"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
	s "github.com/thedevelopnik/netplan/pkg/models"
)

// CreateNetworkMapEndpoint creates a NetworkMap and returns the created value.
// Returns a 400 if it  can't create the struct,
// or a 500 if the db connection or creation fails.
func (h netplanHTTP) CreateNetworkMapEndpoint(c *gin.Context) {
	// get the network map object from the request, or send error
	var nm s.NetworkMap
	if err := c.ShouldBindJSON(&nm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
	}

	// create in the db
	if err := h.svc.CreateNetworkMap(&nm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
	}

	// send full created database object back
	c.JSON(http.StatusCreated, nm)
}

// GetNetworkMapEndpoint gets a NetworkMap struct from a given id.
func (h netplanHTTP) GetNetworkMapEndpoint(c *gin.Context) {
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

	nm, err := h.svc.GetNetworkMap(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
	}

	// return it in the response
	c.JSON(http.StatusOK, nm)
}

// GetAllNetworkMapsEndpoint retrieves all network maps from the database and returns them.
func (h netplanHTTP) GetAllNetworkMapsEndpoint(c *gin.Context) {
	networkMaps, err := h.svc.GetAllNetworkMaps()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errors.WithStack(err),
		})
	}

	c.JSON(http.StatusOK, networkMaps)
}

// UpdateNetworkMapEndpoint updates the name of a
// NetworkMap given an id and name.
func (h netplanHTTP) UpdateNetworkMapEndpoint(c *gin.Context) {
	// get the values to update with off the request
	var nm s.NetworkMap
	if err := c.ShouldBindJSON(&nm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
	}

	update, err := h.svc.UpdateNetworkMap(&nm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
	}

	// send back updated value
	c.JSON(http.StatusOK, update)
}

// DeleteNetworkMapEndpoint deletes a NetworkMap given an id.
func (h netplanHTTP) DeleteNetworkMapEndpoint(c *gin.Context) {
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

	// find db ojbect matching the id
	if err := h.svc.DeleteNetworkMap(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
	}

	// send back a no content response
	c.Status(http.StatusNoContent)
}
