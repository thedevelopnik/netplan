package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	s "github.com/thedevelopnik/netplan/structs"
)

func CreateVPCEndpoint(c *gin.Context) {
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not access database connection",
		})
	}

	var vpc s.VPC
	if err := c.ShouldBindJSON(&vpc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
	}

	if err := db.Create(&vpc).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
	}

	c.JSON(http.StatusCreated, vpc)
}

func UpdateVPCEndpoint(c *gin.Context) {
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not access database connection",
		})
	}

	var vpc s.VPC
	if err := c.ShouldBindJSON(&vpc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
	}

	var update s.VPC
	db.Where("id = ?", vpc.ID).First(&update)
	update.Name = vpc.Name
	db.Save(&update)
	c.JSON(http.StatusOK, update)
}

func DeleteVPCEndpoint(c *gin.Context) {
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

	var vpc s.VPC
	db.Where("id = ?", id).First(&vpc)
	db.Delete(vpc)
	c.Status(http.StatusNoContent)
}
