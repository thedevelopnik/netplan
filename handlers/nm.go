package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	s "github.com/thedevelopnik/netplan/structs"
)

// CreateNetworkMapEndpoint creates a NetworkMap and returns the created value.
// Returns a 400 if it  can't create the struct,
// or a 500 if the db connection or creation fails.
func CreateNetworkMapEndpoint(c *gin.Context) {
	// get the database connection from the context
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not access database connection",
		})
	}

	// get the network map object from the request, or send error
	var nm s.NetworkMap
	if err := c.ShouldBindJSON(&nm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
	}

	// create in the db
	if err := db.Create(&nm).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
	}

	// send full created database object back
	c.JSON(http.StatusCreated, nm)
}

// GetNetworkMapEndpoint gets a NetworkMap struct from a given id.
func GetNetworkMapEndpoint(c *gin.Context) {
	// get the database connection from the context
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not access database connection",
		})
	}

	id, err := convertParamToInt("id", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
	}

	// find the network map in the db
	var nm s.NetworkMap
	if err := db.Where("id = ?", id).First(&nm).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
	}

	// return it in the response
	c.JSON(http.StatusOK, nm)
}

// UpdateNetworkMapEndpoint updates the name of a
// NetworkMap given an id and name.
func UpdateNetworkMapEndpoint(c *gin.Context) {
	// get the database connection from the context
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not access database connection",
		})
	}

	// get the values to update with off the request
	var nm s.NetworkMap
	if err := c.ShouldBindJSON(&nm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
	}

	// find the current one matching the one with updated values
	var update s.NetworkMap
	if err := db.Where("id = ?", nm.ID).First(&update).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
	}

	// update the name
	update.Name = nm.Name

	// save in the db or send error
	if err := db.Save(&update).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
	}

	// send back updated value
	c.JSON(http.StatusOK, update)
}

// DeleteNetworkMapEndpoint deletes a NetworkMap given an id.
func DeleteNetworkMapEndpoint(c *gin.Context) {
	// get the database connection from the context
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not access database connection",
		})
	}

	id, err := convertParamToInt("id", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
	}

	// find db ojbect matching the id
	var nm s.NetworkMap
	if err := db.Where("id = ?", id).First(&nm).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
	}

	// delete the object
	if err := db.Delete(nm).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
	}

	// send back a no content response
	c.Status(http.StatusNoContent)
}
