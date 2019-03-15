package structs

import "github.com/jinzhu/gorm"

// DB Models
type (
	// Subnet represents a subnet model
	Subnet struct {
		gorm.Model
		Name      string `gorm:"not null"`
		Access    string `gorm:"not null"`
		Location  string `gorm:"not null"`
		Provider  string `gorm:"not null"`
		Env       string `gorm:"not null"`
		CidrBlock string `gorm:"not null"`
		VPCID     uint
	}

	// VPC represents a VPC model
	VPC struct {
		gorm.Model
		Name         string `gorm:"not null"`
		Access       string `gorm:"not null"`
		Location     string `gorm:"not null"`
		Provider     string `gorm:"not null"`
		Env          string `gorm:"not null"`
		CidrBlock    string `gorm:"not null"`
		Type         string `gorm:"not null"`
		Subnets      []Subnet
		NetworkMapID uint
	}

	// NetworkMap represents a network map model
	NetworkMap struct {
		gorm.Model
		Name string `gorm:"not null"`
		VPCs []VPC
	}
)
