package db

import (
	"github.com/pkg/errors"
	m "github.com/thedevelopnik/netplan/pkg/models"
)

// VPC actions
func (r npRepo) CreateVPC(vpc *m.VPC) error {
	// create in the db
	if err := r.db.Create(&vpc).Error; err != nil {
		return err
	}
	return nil
}

func (r npRepo) UpdateVPC(vpc *m.VPC) (*m.VPC, error) {
	// find the current one matching the one with updated values
	var update m.VPC
	if err := r.db.Where("id = ?", vpc.ID).First(&update).Error; err != nil {
		return nil, err
	}

	// update the values
	update.Name = vpc.Name
	update.Access = vpc.Access
	update.Location = vpc.Location
	update.Provider = vpc.Provider
	update.Env = vpc.Env
	update.CidrBlock = vpc.CidrBlock
	update.Type = vpc.Type
	update.NetworkMapID = vpc.NetworkMapID

	// save in the db or send error
	if err := r.db.Save(&update).Error; err != nil {
		return nil, err
	}

	return &update, nil
}

func (r npRepo) DeleteVPC(id uint) error {
	// find db ojbect matching the id
	var vpc m.VPC
	if err := r.db.Where("id = ?", id).First(&vpc).Error; err != nil {
		return err
	}

	// delete the object
	if err := r.db.Delete(vpc).Error; err != nil {
		return err
	}

	return nil
}

func (r npRepo) GetVPCsByNetworkMapID(id uint) ([]m.VPC, error) {
	var vpcs []m.VPC
	if err := r.db.Where("network_map_id = ?", id).Find(&vpcs).Error; err != nil {
		return nil, errors.Wrap(err, "repository could not find vpcs matching the network map id in the db")
	}
	return vpcs, nil
}
