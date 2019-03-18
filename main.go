package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	database "github.com/thedevelopnik/netplan/pkg/db"
	"github.com/thedevelopnik/netplan/pkg/service"
	"github.com/thedevelopnik/netplan/pkg/transport"
)

// need to:
// think critically about error handling from db up to response, especially in instances with multiple possible errors
// test the new service package
// retest the transport/handler package
// have better logging, probably logrus
// do better config, especially of db values

func main() {
	db, err := database.Conn()
	if err != nil {
		log.Fatalln(err)
	}

	repo := database.New(db)
	svc := service.New(repo)
	t := transport.New(svc)

	r := gin.Default()
	r.Use(dbMiddleware(db))
	r.StaticFile("/app.js", "./dist/app.js")
	r.StaticFile("/about.js", "./dist/about.js")
	r.StaticFile("/", "./dist/index.html")
	v1 := r.Group("/v1")
	{
		// NetworkMap Endpoints
		v1.POST("/networkmap", t.CreateNetworkMapEndpoint)
		v1.GET("/networkmap", t.GetAllNetworkMapsEndpoint)
		v1.GET("/networkmap/:id", t.GetNetworkMapEndpoint)
		v1.PUT("/networkmap", t.UpdateNetworkMapEndpoint)
		v1.DELETE("/networkmap/:id", t.DeleteNetworkMapEndpoint)

		// VPC Endpoints
		v1.POST("/networkmap/:nmid/vpc", t.CreateVPCEndpoint)
		v1.PUT("/networkmap/:nmid/vpc/:id", t.UpdateVPCEndpoint)
		v1.DELETE("/networkmap/:nmid/vpc/:id", t.DeleteVPCEndpoint)

		// Subnet Endpoints
		v1.POST("/networkmap/:nmid/vpc/:vpcid/subnet", t.CreateSubnetEndpoint)
		v1.PUT("/networkmap/:nmid/vpc/:pvcid/subnet/:id", t.UpdateSubnetEndpoint)
		v1.DELETE("/networkmap/:nmid/vpc/:vpcid/subnet/:id", t.DeleteSubnetEndpoint)
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
