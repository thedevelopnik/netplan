package db

import (
	"github.com/jinzhu/gorm"
	m "github.com/thedevelopnik/netplan/pkg/models"
)

type NetplanRepository interface {
	CreateNetworkMap(*m.NetworkMap) error
	GetNetworkMap(uint) (*m.NetworkMap, error)
	GetAllNetworkMaps() ([]m.NetworkMap, error)
	UpdateNetworkMap(*m.NetworkMap) (*m.NetworkMap, error)
	DeleteNetworkMap(uint) error
	CreateVPC(*m.VPC) error
	UpdateVPC(*m.VPC) (*m.VPC, error)
	DeleteVPC(uint) error
	CreateSubnet(*m.Subnet) error
	UpdateSubnet(*m.Subnet) (*m.Subnet, error)
	DeleteSubnet(uint) error
	GetVPCsByNetworkMapID(uint) ([]m.VPC, error)
	GetSubnetsByVPCID(uint) ([]m.Subnet, error)
}

func New(db *gorm.DB) NetplanRepository {
	return &npRepo{
		db: db,
	}
}

type npRepo struct {
	db *gorm.DB
}
