package service

import (
	"github.com/pkg/errors"
	"github.com/thedevelopnik/netplan/pkg/db"
	m "github.com/thedevelopnik/netplan/pkg/models"
	"github.com/thedevelopnik/netplan/pkg/networks"
)

type NetPlan interface {
	CreateNetworkMap(*m.NetworkMap) error
	GetNetworkMap(uint) (*m.NetworkMap, error)
	UpdateNetworkMap(*m.NetworkMap) (*m.NetworkMap, error)
	DeleteNetworkMap(uint) error
	CreateVPC(*m.VPC) error
	UpdateVPC(*m.VPC) (*m.VPC, error)
	DeleteVPC(uint) error
	CreateSubnet(*m.Subnet) error
	UpdateSubnet(*m.Subnet) (*m.Subnet, error)
	DeleteSubnet(uint) error
	checkVPCOverlap(m.VPC, []m.VPC) (bool, error)
	checkSubnetOverlap(m.Subnet, []m.Subnet) (bool, error)
}

func New(repo db.NetplanRepository) NetPlan {
	return netplan{
		repo: repo,
	}
}

type netplan struct {
	repo db.NetplanRepository
}

func (svc netplan) checkVPCOverlap(vpc m.VPC, existingVPCs []m.VPC) (bool, error) {
	newVPCNetwork, err := networks.New(vpc.CidrBlock)
	if err != nil {
		return false, errors.Wrap(err, "could not create network out of the vpc cidrblock")
	}
	for _, eVPC := range existingVPCs {
		network, err := networks.New(eVPC.CidrBlock)
		if err != nil {
			return false, errors.Wrap(err, "could not create network out of existing vpc for network map")
		}
		if network.Cidr.Contains(*newVPCNetwork.First) || network.Cidr.Contains(*newVPCNetwork.Last) {
			return true, nil
		}
		return false, nil
	}
	return false, nil
}

func (svc netplan) checkSubnetOverlap(subnet m.Subnet, existingSubnets []m.Subnet) (bool, error) {
	newSubNetwork, err := networks.New(subnet.CidrBlock)
	if err != nil {
		return false, errors.Wrap(err, "could not create network out of the vpc cidrblock")
	}
	for _, eSubnet := range existingSubnets {
		network, err := networks.New(eSubnet.CidrBlock)
		if err != nil {
			return false, errors.Wrap(err, "could not create network out of existing vpc for network map")
		}
		if network.Cidr.Contains(*newSubNetwork.First) || network.Cidr.Contains(*newSubNetwork.Last) {
			return true, nil
		}
		return false, nil
	}
	return false, nil
}
