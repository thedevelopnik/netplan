package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/thedevelopnik/netplan/db"
	_ "github.com/thedevelopnik/netplan/db"
	h "github.com/thedevelopnik/netplan/handlers"
)

func main() {
	db := db.Conn()
	defer db.Close()

	r := gin.Default()
	r.Use(dbMiddleware(db))
	r.StaticFile("/app.js", "./dist/app.js")
	r.StaticFile("/about.js", "./dist/about.js")
	r.StaticFile("/", "./dist/index.html")
	v1 := r.Group("/v1")
	{
		// NetworkMap Endpoints
		v1.POST("/networkmap", h.CreateNetworkMapEndpoint)
		v1.GET("/networkmap/:id", h.GetNetworkMapEndpoint)
		v1.PUT("/networkmap", h.UpdateNetworkMapEndpoint)
		v1.DELETE("/networkmap/:id", h.DeleteNetworkMapEndpoint)

		// VPC Endpoints
		v1.POST("/networkmap/:nmid/vpc", h.CreateVPCEndpoint)
		v1.PUT("/networkmap/:nmid/vpc/:id", h.UpdateVPCEndpoint)
		v1.DELETE("/networkmap/:nmid/vpc/:id", h.DeleteVPCEndpoint)

		// Subnet Endpoints
		v1.POST("/networkmap/:nmid/vpc/:vpcid/subnet", h.CreateSubnetEndpoint)
		v1.PUT("/networkmap/:nmid/vpc/:pvcid/subnet/:id", h.UpdateSubnetEndpoint)
		v1.DELETE("/networkmap/:nmid/vpc/:vpcid/subnet/:id", h.DeleteSubnetEndpoint)
	}
	r.Run() // listen and serve on 0.0.0.0:8080
}

func dbMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
