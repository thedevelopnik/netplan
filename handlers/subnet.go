package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	s "github.com/thedevelopnik/netplan/structs"
)

func CreateSubnetEndpoint(c *gin.Context) {
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not access database connection",
		})
	}

	var sn s.Subnet
	if err := c.ShouldBindJSON(&sn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
	}

	if err := db.Create(&sn).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
	}

	c.JSON(http.StatusCreated, sn)
}

func UpdateSubnetEndpoint(c *gin.Context) {
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not access database connection",
		})
	}

	var sn s.Subnet
	if err := c.ShouldBindJSON(&sn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
	}

	var update s.Subnet
	db.Where("id = ?", sn.ID).First(&update)
	update.Name = sn.Name
	db.Save(&update)
	c.JSON(http.StatusOK, update)
}

func DeleteSubnetEndpoint(c *gin.Context) {
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

	var sn s.Subnet
	db.Where("id = ?", id).First(&sn)
	db.Delete(sn)
	c.Status(http.StatusNoContent)
}
