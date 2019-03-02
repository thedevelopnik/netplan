package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	s "github.com/thedevelopnik/netplan/structs"
)

func Conn() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=6666 user=netplan dbname=netplan password=netplan sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&s.Subnet{}, &s.VPC{}, &s.NetworkMap{})
	return db
}
