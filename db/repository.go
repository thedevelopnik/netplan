package db

import (
	"github.com/jinzhu/gorm"
	s "github.com/thedevelopnik/netplan/structs"
)

type NetplanRepository interface {
	CreateNetworkMap(*s.NetworkMap) error
	GetNetworkMap(uint) (*s.NetworkMap, error)
	UpdateNetworkMap(*s.NetworkMap) (*s.NetworkMap, error)
	DeleteNetworkMap(uint) error
	CreateVPC(*s.VPC) error
	UpdateVPC(*s.VPC) (*s.VPC, error)
	DeleteVPC(uint) error
	CreateSubnet(*s.Subnet) error
	UpdateSubnet(*s.Subnet) (*s.Subnet, error)
	DeleteSubnet(uint) error
}

func New(db *gorm.DB) NetplanRepository {
	return &npRepo{
		db: db,
	}
}

type npRepo struct {
	db *gorm.DB
}
