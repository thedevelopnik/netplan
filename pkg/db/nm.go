package db

import (
	"github.com/pkg/errors"
	m "github.com/thedevelopnik/netplan/pkg/models"
)

// NetworkMap actions
func (r npRepo) CreateNetworkMap(nm *m.NetworkMap) error {
	// create in the db
	if err := r.db.Create(&nm).Error; err != nil {
		return err
	}

	return nil
}

func (r npRepo) GetNetworkMap(id uint) (*m.NetworkMap, error) {
	var nm m.NetworkMap
	if err := r.db.Set("gorm:auto_preload", true).Where("id = ?", id).First(&nm).Error; err != nil {
		return nil, err
	}

	return &nm, nil
}

func (r npRepo) GetAllNetworkMaps() ([]m.NetworkMap, error) {
	var networkMaps []m.NetworkMap
	if err := r.db.Find(&networkMaps).Error; err != nil {
		return nil, errors.Wrap(err, "could not retrieve list network maps")
	}
	return networkMaps, nil
}

func (r npRepo) UpdateNetworkMap(nm *m.NetworkMap) (*m.NetworkMap, error) {
	// find the current one matching the one with updated values
	var update m.NetworkMap
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
	var nm m.NetworkMap
	if err := r.db.Where("id = ?", id).First(&nm).Error; err != nil {
		return err
	}

	// delete the object
	if err := r.db.Delete(nm).Error; err != nil {
		return err
	}

	return nil
}
