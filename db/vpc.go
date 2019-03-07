package db

import (
	s "github.com/thedevelopnik/netplan/structs"
)

// VPC actions
func (r npRepo) CreateVPC(vpc *s.VPC) error {
	// create in the db
	if err := r.db.Create(&vpc).Error; err != nil {
		return err
	}
	return nil
}

func (r npRepo) UpdateVPC(vpc *s.VPC) (*s.VPC, error) {
	// find the current one matching the one with updated values
	var update s.VPC
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

func (r npRepo) DeleteVPC(id int) error {
	// find db ojbect matching the id
	var vpc s.VPC
	if err := r.db.Where("id = ?", id).First(&vpc).Error; err != nil {
		return err
	}

	// delete the object
	if err := r.db.Delete(vpc).Error; err != nil {
		return err
	}

	return nil
}
