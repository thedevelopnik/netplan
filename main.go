package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	s "github.com/thedevelopnik/netplan/structs"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("postgres", "host=localhost port=6666 user=netplan dbname=netplan password=netplan sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&s.Subnet{}, &s.VPC{}, &s.NetworkMap{})
}

func main() {
	defer db.Close()

	r := gin.Default()
	r.StaticFile("/app.js", "./dist/app.js")
	r.StaticFile("/about.js", "./dist/about.js")
	r.StaticFile("/", "./dist/index.html")
	v1 := r.Group("/v1")
	{
		v1.POST("/networkmap", createNetworkMapEndpoint)
		v1.GET("/networkmap/:id", getNetworkMapEndpoint)
		v1.PUT("/networkmap/:id", updateNetworkMapEndpoint)
		v1.DELETE("/networkmap/:id", deleteNetworkMapEndpoint)

		v1.POST("/networkmap/:id/vpc", createVPCEndpoint)
		v1.PUT("/networkmap/:id/vpc/:id", updateVPCEndpoint)
		v1.DELETE("/networkmap/:id/vpc/:id", deleteVPCEndpoint)

		v1.POST("/network/:id/vpc/:id/subnet", createSubnetEndpoint)
		v1.PUT("/network/:id/vpc/:id/subnet/:id", updateSubnetEndpoint)
		v1.DELETE("/network/:id/vpc/:id/subnet/:id", deleteSubnetEndpoint)
	}
	r.Run() // listen and serve on 0.0.0.0:8080
}

func createNetworkMapEndpoint(c *gin.Context) {
	var req s.CreateNetworkMapRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
	}
	nm := s.NetworkMap{Name: req.Name}
	db.Create(&nm)
	c.JSON(http.StatusCreated, nm)
}

func getNetworkMapEndpoint(c *gin.Context) {
	sid := c.Param("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
	}
	var nm s.NetworkMap
	db.Where("id = ?", id).First(&nm)
	c.JSON(http.StatusOK, nm)
}

func updateNetworkMapEndpoint(c *gin.Context) {}

func deleteNetworkMapEndpoint(c *gin.Context) {}

func createVPCEndpoint(c *gin.Context) {}

func updateVPCEndpoint(c *gin.Context) {}

func deleteVPCEndpoint(c *gin.Context) {}

func createSubnetEndpoint(c *gin.Context) {}

func updateSubnetEndpoint(c *gin.Context) {}

func deleteSubnetEndpoint(c *gin.Context) {}
