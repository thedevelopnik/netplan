package service

import (
	"fmt"

	"github.com/pkg/errors"
	m "github.com/thedevelopnik/netplan/pkg/models"
)

func (svc netplan) CreateSubnet(subnet *m.Subnet) error {
	subnets, err := svc.repo.GetSubnetsByVPCID(subnet.VPCID)
	if err != nil {
		fmt.Println("no existing subnets found for vpc")
	}

	overlap, err := svc.checkSubnetOverlap(*subnet, subnets)
	if err != nil {
		return errors.Wrap(err, "error checking overlap with other subnets in vpc")
	}

	if overlap {
		return errors.New("could not create subnet because of overlap with other subnets in vpc")
	}

	if err := svc.repo.CreateSubnet(subnet); err != nil {
		return errors.Wrap(err, "netplan service could not create subnet in database")
	}

	return nil
}

func (svc netplan) UpdateSubnet(sn *m.Subnet) (*m.Subnet, error) {
	update, err := svc.repo.UpdateSubnet(sn)
	if err != nil {
		return nil, errors.Wrap(err, "could not update subnet")
	}
	return update, nil
}

func (svc netplan) DeleteSubnet(id uint) error {
	if err := svc.repo.DeleteSubnet(id); err != nil {
		return errors.Wrap(err, "could not delete subnet")
	}
	return nil
}
