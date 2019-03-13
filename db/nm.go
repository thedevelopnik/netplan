package db

import (
	s "github.com/thedevelopnik/netplan/structs"
)

// NetworkMap actions
func (r npRepo) CreateNetworkMap(nm *s.NetworkMap) error {
	// create in the db
	if err := r.db.Create(&nm).Error; err != nil {
		return err
	}

	return nil
}

func (r npRepo) GetNetworkMap(id uint) (*s.NetworkMap, error) {
	var nm s.NetworkMap
	if err := r.db.Where("id = ?", id).First(&nm).Error; err != nil {
		return nil, err
	}

	return &nm, nil
}

func (r npRepo) UpdateNetworkMap(nm *s.NetworkMap) (*s.NetworkMap, error) {
	// find the current one matching the one with updated values
	var update s.NetworkMap
	if err := r.db.Where("id = ?", nm.ID).First(&update).Error; err != nil {
		return nil, err
	}

	// update the name
	update.Name = nm.Name

	// save in the db or send error
	if err := r.db.Save(&update).Error; err != nil {
		return nil, err
	}

	return &update, nil
}

func (r npRepo) DeleteNetworkMap(id uint) error {
	// find db ojbect matching the id
	var nm s.NetworkMap
	if err := r.db.Where("id = ?", id).First(&nm).Error; err != nil {
		return err
	}

	// delete the object
	if err := r.db.Delete(nm).Error; err != nil {
		return err
	}

	return nil
}
