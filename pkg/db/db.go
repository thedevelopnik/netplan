package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	s "github.com/thedevelopnik/netplan/pkg/models"
)

func Conn() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=localhost port=6666 user=netplan dbname=netplan password=netplan sslmode=disable")
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&s.Subnet{}, &s.VPC{}, &s.NetworkMap{})
	return db, nil
}
