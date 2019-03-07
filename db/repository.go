package db

import (
	"github.com/jinzhu/gorm"
	s "github.com/thedevelopnik/netplan/structs"
)

type NetplanRepository interface {
	CreateNetworkMap(*s.NetworkMap) error
	GetNetworkMap(int) (*s.NetworkMap, error)
	UpdateNetworkMap(*s.NetworkMap) (*s.NetworkMap, error)
	DeleteNetworkMap(int) error
	CreateVPC(*s.VPC) (*s.VPC, error)
	UpdateVPC(*s.VPC) (*s.VPC, error)
	DeleteVPC(int) error
	CreateSubnet(*s.Subnet) (*s.Subnet, error)
	UpdateSubnet(*s.Subnet) (*s.Subnet, error)
	DeleteSubnet(int) error
}

func New(db *gorm.DB) NetplanRepository {
	return &npRepo{
		db: db,
	}
}

type npRepo struct {
	db *gorm.DB
}
