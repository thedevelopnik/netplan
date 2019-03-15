package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevelopnik/netplan/pkg/service"
)

type NetplanHTTP interface {
	CreateNetworkMapEndpoint(*gin.Context)
	GetNetworkMapEndpoint(*gin.Context)
	UpdateNetworkMapEndpoint(*gin.Context)
	DeleteNetworkMapEndpoint(*gin.Context)
	CreateVPCEndpoint(*gin.Context)
	UpdateVPCEndpoint(*gin.Context)
	DeleteVPCEndpoint(*gin.Context)
	CreateSubnetEndpoint(*gin.Context)
	UpdateSubnetEndpoint(*gin.Context)
	DeleteSubnetEndpoint(*gin.Context)
}

func New(svc service.NetPlan) NetplanHTTP {
	return netplanHTTP{
		svc: svc,
	}
}

type netplanHTTP struct {
	svc service.NetPlan
}
