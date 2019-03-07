package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"

	database "github.com/thedevelopnik/netplan/db"
	h "github.com/thedevelopnik/netplan/handlers"
)

func main() {
	db := database.Conn()
	repo := database.New(db)
	svc := h.New(repo)

	r := gin.Default()
	r.Use(dbMiddleware(db))
	r.StaticFile("/app.js", "./dist/app.js")
	r.StaticFile("/about.js", "./dist/about.js")
	r.StaticFile("/", "./dist/index.html")
	v1 := r.Group("/v1")
	{
		// NetworkMap Endpoints
		v1.POST("/networkmap", svc.CreateNetworkMapEndpoint)
		v1.GET("/networkmap/:id", svc.GetNetworkMapEndpoint)
		v1.PUT("/networkmap", svc.UpdateNetworkMapEndpoint)
		v1.DELETE("/networkmap/:id", svc.DeleteNetworkMapEndpoint)

		// VPC Endpoints
		v1.POST("/networkmap/:nmid/vpc", svc.CreateVPCEndpoint)
		v1.PUT("/networkmap/:nmid/vpc/:id", svc.UpdateVPCEndpoint)
		v1.DELETE("/networkmap/:nmid/vpc/:id", svc.DeleteVPCEndpoint)

		// Subnet Endpoints
		v1.POST("/networkmap/:nmid/vpc/:vpcid/subnet", svc.CreateSubnetEndpoint)
		v1.PUT("/networkmap/:nmid/vpc/:pvcid/subnet/:id", svc.UpdateSubnetEndpoint)
		v1.DELETE("/networkmap/:nmid/vpc/:vpcid/subnet/:id", svc.DeleteSubnetEndpoint)
	}
	if err := r.Run(); err != nil {
		log.Fatalln(err)
	} // listen and serve on 0.0.0.0:8080
}

func dbMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
