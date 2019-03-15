package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevelopnik/netplan/pkg/db"
)

type NetplanService interface {
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

func New(repo db.NetplanRepository) NetplanService {
	return netplanService{
		repo: repo,
	}
}

type netplanService struct {
	repo db.NetplanRepository
}
