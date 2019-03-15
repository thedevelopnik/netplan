package db

import (
	"strings"
	"testing"

	s "github.com/thedevelopnik/netplan/structs"
)

func TestCreateSubnet(t *testing.T) {
	sn := s.Subnet{
		Name:      "create-test-subnet",
		Access:    "public",
		Location:  "us-east4",
		Provider:  "GCP",
		Env:       "dev",
		CidrBlock: "192.168.0.0/16",
		VPCID:     1,
	}

	conn, err := Conn()
	if err != nil {
		t.Error(err)
	}
	repo := New(conn)

	err = repo.CreateSubnet(&sn)
	if err != nil {
		t.Error(err)
	}

	if sn.ID == 0 {
		t.Error("subnet did not have valid ID")
	}
	if strings.Compare(sn.Name, "create-test-subnet") != 0 {
		t.Error("subnet name was not correct")
	}

	err = conn.Close()
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateSubnet(t *testing.T) {
	sn := s.Subnet{
		Name:      "original-test-subnet",
		Access:    "public",
		Location:  "us-east4",
		Provider:  "GCP",
		Env:       "dev",
		CidrBlock: "192.168.0.0/16",
		VPCID:     1,
	}

	conn, err := Conn()
	if err != nil {
		t.Error(err)
	}
	repo := New(conn)

	err = repo.CreateSubnet(&sn)
	if err != nil {
		t.Error(err)
	}
	updater := s.Subnet{}
	updater.ID = sn.ID
	updater.Name = "update-test-subnet"
	updater.CidrBlock = "10.10.0.0/26"

	usn, err := repo.UpdateSubnet(&updater)

	if strings.Compare(usn.Name, "update-test-subnet") != 0 {
		t.Error("updated subnet does not have correct name")
	}

	if strings.Compare(usn.CidrBlock, "10.10.0.0/26") != 0 {
		t.Error("updated subnet does not have correct cidr block")
	}

	if usn.ID != sn.ID {
		t.Error("updated subnet does not represent the correct db record")
	}

	err = conn.Close()
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteSubnet(t *testing.T) {
	sn := s.Subnet{
		Name:      "delete-test-subnet",
		Access:    "public",
		Location:  "us-east4",
		Provider:  "GCP",
		Env:       "dev",
		CidrBlock: "192.168.0.0/16",
		VPCID:     1,
	}

	conn, err := Conn()
	if err != nil {
		t.Error(err)
	}
	repo := New(conn)

	err = repo.CreateSubnet(&sn)
	if err != nil {
		t.Error(err)
	}
	err = repo.DeleteSubnet(sn.ID)
	if err != nil {
		t.Error(err)
	}

	var dsn s.Subnet
	err = conn.Where("id = ?", sn.ID).First(&dsn).Error
	if err == nil {
		t.Error("found a subnet that should have been deleted")
	}

	err = conn.Close()
	if err != nil {
		t.Error(err)
	}
}
