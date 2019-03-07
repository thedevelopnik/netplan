package db

import (
	s "github.com/thedevelopnik/netplan/structs"
)

// Subnet actions
func (r npRepo) CreateSubnet(sn *s.Subnet) (*s.Subnet, error) {
	// create in the db
	if err := db.Create(&sn).Error; err != nil {
		return nil, err
	}

	return sn, nil
}

func (r npRepo) UpdateSubnet(sn *s.Subnet) (*s.Subnet, error) {
	var update s.Subnet
	if err := db.Where("id = ?", sn.ID).First(&update).Error; err != nil {
		return nil, err
	}

	// update the values
	update.Name = sn.Name
	update.Access = sn.Access
	update.Location = sn.Location
	update.Provider = sn.Provider
	update.Env = sn.Env
	update.CidrBlock = sn.CidrBlock
	update.Type = sn.Type
	update.VPCID = sn.VPCID

	// save in the db or send error
	if err := db.Save(&update).Error; err != nil {
		return nil, err
	}

	return update, nil
}

func (r npRepo) DeleteSubnet(id int) error {
	// find db ojbect matching the id
	var sn s.Subnet
	if err := db.Where("id = ?", id).First(&sn).Error; err != nil {
		return err
	}

	// delete the object
	if err := db.Delete(sn).Error; err != nil {
		return err
	}
}
