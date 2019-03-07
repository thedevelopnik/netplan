package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	s "github.com/thedevelopnik/netplan/structs"
)

// CreateNetworkMapEndpoint creates a NetworkMap and returns the created value.
// Returns a 400 if it  can't create the struct,
// or a 500 if the db connection or creation fails.
func CreateNetworkMapEndpoint(c *gin.Context) {
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not access database connection",
		})
	}

	var nm s.NetworkMap
	if err := c.ShouldBindJSON(&nm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
	}

	if err := db.Create(&nm).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
	}

	c.JSON(http.StatusCreated, nm)
}

// GetNetworkMapEndpoint gets a NetworkMap struct from a given id.
func GetNetworkMapEndpoint(c *gin.Context) {
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not access database connection",
		})
	}

	sid := c.Param("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
	}

	var nm s.NetworkMap
	if err := db.Where("id = ?", id).First(&nm).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
	}

	c.JSON(http.StatusOK, nm)
}

// UpdateNetworkMapEndpoint updates the name of a
// NetworkMap given an id and name.
func UpdateNetworkMapEndpoint(c *gin.Context) {
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not access database connection",
		})
	}

	var nm s.NetworkMap
	if err := c.ShouldBindJSON(&nm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
	}

	var update s.NetworkMap
	db.Where("id = ?", nm.ID).First(&update)
	update.Name = nm.Name
	db.Save(&update)
	c.JSON(http.StatusOK, update)
}

// DeleteNetworkMapEndpoint deletes a NetworkMap given an id.
func DeleteNetworkMapEndpoint(c *gin.Context) {
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not access database connection",
		})
	}

	sid := c.Param("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
	}

	var nm s.NetworkMap
	db.Where("id = ?", id).First(&nm)
	db.Delete(nm)
	c.Status(http.StatusNoContent)
}
