package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	s "github.com/thedevelopnik/netplan/structs"
)

// CreateVPCEndpoint creates a VPC and returns the created value.
// Returns a 400 if it  can't create the struct,
// or a 500 if the db connection or creation fails.
func CreateVPCEndpoint(c *gin.Context) {
	// get the database connection from the context
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not access database connection",
		})
	}

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
	if err := db.Create(&vpc).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
	}

	// send full created database object back
	c.JSON(http.StatusCreated, vpc)
}

// UpdateVPCEndpoint updates the name of a
// VPC given an id and name.
func UpdateVPCEndpoint(c *gin.Context) {
	// get the database connection from the context
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not access database connection",
		})
	}

	// get the values to update with off the request
	var vpc s.VPC
	if err := c.ShouldBindJSON(&vpc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
	}

	// find the current one matching the one with updated values
	var update s.VPC
	if err := db.Where("id = ?", vpc.ID).First(&update).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
	}

	// update the name
	update.Name = vpc.Name

	// save in the db or send error
	if err := db.Save(&update).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
	}

	// send back updated value
	c.JSON(http.StatusOK, update)
}

// DeleteVPCEndpoint deletes a VPC given an id.
func DeleteVPCEndpoint(c *gin.Context) {
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
	var vpc s.VPC
	if err := db.Where("id = ?", id).First(&vpc).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
	}

	// delete the object
	if err := db.Delete(vpc).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
	}

	// send back a no content response
	c.Status(http.StatusNoContent)
}
