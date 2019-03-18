package service

import (
	"fmt"

	"github.com/pkg/errors"
	m "github.com/thedevelopnik/netplan/pkg/models"
)

func (svc netplan) CreateNetworkMap(nm *m.NetworkMap) error {
	if err := svc.repo.CreateNetworkMap(nm); err != nil {
		return errors.Wrap(err, "netplan service could not create network map")
	}
	return nil
}

func (svc netplan) GetNetworkMap(id uint) (*m.NetworkMap, error) {
	nm, err := svc.repo.GetNetworkMap(id)
	if err != nil {
		return nil, errors.Wrap(err, "netplan service could not get network map")
	}
	return nm, nil
}

func (svc netplan) GetAllNetworkMaps() ([]m.NetworkMap, error) {
	networkMaps, err := svc.repo.GetAllNetworkMaps()
	if err != nil {
		return nil, errors.Wrap(err, "netplan service could not get list of network maps")
	}
	return networkMaps, nil
}

func (svc netplan) UpdateNetworkMap(nm *m.NetworkMap) (*m.NetworkMap, error) {
	update, err := svc.repo.UpdateNetworkMap(nm)
	if err != nil {
		return nil, errors.Wrap(err, "netplan service could not update network map")
	}
	return update, nil
}

func (svc netplan) DeleteNetworkMap(id uint) error {
	// get network map with all vpcs and subnets
	var allErrors []error
	nm, err := svc.repo.GetNetworkMap(id)
	if err != nil {
		return errors.Wrap(err, "netplan service could not find network map to delete it")
	}
	// delete all associated vpcs and subnets
	vpcs := nm.VPCs
	for _, vpc := range vpcs {
		for _, subnet := range vpc.Subnets {
			err = svc.repo.DeleteSubnet(subnet.ID)
			if err != nil {
				allErrors = append(allErrors, err)
			}
		}
		err = svc.repo.DeleteVPC(vpc.ID)
		if err != nil {
			allErrors = append(allErrors, err)
		}
	}

	// delete network map
	if err := svc.repo.DeleteNetworkMap(id); err != nil {
		allErrors = append(allErrors, err)
	}

	if len(allErrors) > 0 {
		for _, err := range allErrors {
			fmt.Println(err)
		}
		return errors.New("could not delete network map")
	}

	return nil
}
