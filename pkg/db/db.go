package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/thedevelopnik/netplan/pkg/config"
	s "github.com/thedevelopnik/netplan/pkg/models"
)

func Conn(conf config.DBConfig) (*gorm.DB, error) {
	args := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		conf.Host,
		conf.Port,
		conf.User,
		conf.DBName,
		conf.Password,
		conf.SSLMode,
	)
	fmt.Println(args)
	db, err := gorm.Open("postgres", args)
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&s.Subnet{}, &s.VPC{}, &s.NetworkMap{})
	return db, nil
}
