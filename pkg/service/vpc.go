package service

import (
	"fmt"

	"github.com/pkg/errors"
	m "github.com/thedevelopnik/netplan/pkg/models"
)

func (svc netplan) CreateVPC(vpc *m.VPC) error {
	vpcs, err := svc.repo.GetVPCsByNetworkMapID(vpc.NetworkMapID)
	if err != nil {
		fmt.Println("no existing vpcs found for network map id")
	}

	overlap, err := svc.checkVPCOverlap(*vpc, vpcs)
	if err != nil {
		return errors.Wrap(err, "error checking overlap with other vpcs in network map")
	}

	if overlap {
		return errors.New("could not create vpc because of overlap with other vpcs in network map")
	}

	if err := svc.repo.CreateVPC(vpc); err != nil {
		return errors.Wrap(err, "netplan service could not create vpc in database")
	}

	return nil
}

func (svc netplan) UpdateVPC(vpc *m.VPC) (*m.VPC, error) {
	// save in the db or send error
	update, err := svc.repo.UpdateVPC(vpc)
	if err != nil {
		return nil, errors.Wrap(err, "netplan service could not update vpc")
	}
	return update, nil
}

func (svc netplan) DeleteVPC(id uint) error {
	// delete all associated subnets
	subnets, err := svc.repo.GetSubnetsByVPCID(id)
	if err != nil {
		return errors.Wrap(err, "could not find subnets to delete for this vpc")
	}

	for _, subnet := range subnets {
		err := svc.repo.DeleteSubnet(subnet.ID)
		if err != nil {
			return errors.Wrap(err, "could not delete subnet in vpc")
		}
	}

	// delete the vpc
	if err := svc.repo.DeleteVPC(id); err != nil {
		return errors.Wrap(err, "netplan service could not delete vpc")
	}

	return nil
}
