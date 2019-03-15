package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	database "github.com/thedevelopnik/netplan/pkg/db"
	h "github.com/thedevelopnik/netplan/pkg/handlers"
	svc "github.com/thedevelopnik/netplan/pkg/service"
)

func main() {
	db, err := database.Conn()
	if err != nil {
		log.Fatalln(err)
	}

	repo := database.New(db)
	svc := svc.New(repo)
	handler := h.New(svc)

	r := gin.Default()
	r.Use(dbMiddleware(db))
	r.StaticFile("/app.js", "./dist/app.js")
	r.StaticFile("/about.js", "./dist/about.js")
	r.StaticFile("/", "./dist/index.html")
	v1 := r.Group("/v1")
	{
		// NetworkMap Endpoints
		v1.POST("/networkmap", handler.CreateNetworkMapEndpoint)
		v1.GET("/networkmap/:id", handler.GetNetworkMapEndpoint)
		v1.PUT("/networkmap", handler.UpdateNetworkMapEndpoint)
		v1.DELETE("/networkmap/:id", handler.DeleteNetworkMapEndpoint)

		// VPC Endpoints
		v1.POST("/networkmap/:nmid/vpc", handler.CreateVPCEndpoint)
		v1.PUT("/networkmap/:nmid/vpc/:id", handler.UpdateVPCEndpoint)
		v1.DELETE("/networkmap/:nmid/vpc/:id", handler.DeleteVPCEndpoint)

		// Subnet Endpoints
		v1.POST("/networkmap/:nmid/vpc/:vpcid/subnet", handler.CreateSubnetEndpoint)
		v1.PUT("/networkmap/:nmid/vpc/:pvcid/subnet/:id", handler.UpdateSubnetEndpoint)
		v1.DELETE("/networkmap/:nmid/vpc/:vpcid/subnet/:id", handler.DeleteSubnetEndpoint)
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
