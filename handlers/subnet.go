package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	s "github.com/thedevelopnik/netplan/structs"
)

// CreateSubnetEndpoint creates a Subnet and returns the created value.
// Returns a 400 if it  can't create the struct,
// or a 500 if the db connection or creation fails.
func CreateSubnetEndpoint(c *gin.Context) {
	// get the database connection from the context
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not access database connection",
		})
	}

	vpcID, err := convertParamToInt("vpcid", c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
	}

	// get the network map object from the request, or send error
	var sn s.Subnet
	if err := c.ShouldBindJSON(&sn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
	}

	sn.VPCID = vpcID

	// create in the db
	if err := db.Create(&sn).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
	}

	// send full created database object back
	c.JSON(http.StatusCreated, sn)
}

// UpdateSubnetEndpoint updates the name of a
// Subnet given an id and name.
func UpdateSubnetEndpoint(c *gin.Context) {
	// get the database connection from the context
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not access database connection",
		})
	}

	// get the values to update with off the request
	var sn s.Subnet
	if err := c.ShouldBindJSON(&sn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
	}

	// find the current one matching the one with updated values
	var update s.Subnet
	if err := db.Where("id = ?", sn.ID).First(&update).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
	}

	// update the name
	update.Name = sn.Name

	// save in the db or send error
	if err := db.Save(&update).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
	}

	// send back updated value
	c.JSON(http.StatusOK, update)
}

// DeleteSubnetEndpoint deletes a Subnet given an id.
func DeleteSubnetEndpoint(c *gin.Context) {
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
	var sn s.Subnet
	if err := db.Where("id = ?", id).First(&sn).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
	}

	// delete the object
	if err := db.Delete(sn).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
	}

	// send back a no content response
	c.Status(http.StatusNoContent)
}
