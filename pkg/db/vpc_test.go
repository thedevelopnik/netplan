package db

import (
	"strings"
	"testing"

	s "github.com/thedevelopnik/netplan/pkg/models"
)

func TestCreateVPC(t *testing.T) {
	vpc := s.VPC{
		Name:         "create-test-vpc",
		Access:       "public",
		Location:     "us-east4",
		Provider:     "GCP",
		Env:          "dev",
		CidrBlock:    "192.168.0.0/16",
		Type:         "vpc",
		NetworkMapID: 1,
	}

	conn, err := Conn()
	if err != nil {
		t.Error(err)
	}
	repo := New(conn)

	err = repo.CreateVPC(&vpc)
	if err != nil {
		t.Error(err)
	}

	if vpc.ID == 0 {
		t.Error("network map did not have valid ID")
	}
	if strings.Compare(vpc.Name, "create-test-vpc") != 0 {
		t.Error("vpc name was not correct")
	}

	err = conn.Close()
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateVPC(t *testing.T) {
	vpc := s.VPC{
		Name:         "original-test-vpc",
		Access:       "public",
		Location:     "us-east4",
		Provider:     "GCP",
		Env:          "dev",
		CidrBlock:    "192.168.0.0/16",
		Type:         "vpc",
		NetworkMapID: 1,
	}

	conn, err := Conn()
	if err != nil {
		t.Error(err)
	}
	repo := New(conn)

	err = repo.CreateVPC(&vpc)
	if err != nil {
		t.Error(err)
	}
	updater := s.VPC{}
	updater.ID = vpc.ID
	updater.Name = "update-test-vpc"
	updater.CidrBlock = "10.10.0.0/26"

	uvpc, err := repo.UpdateVPC(&updater)

	if strings.Compare(uvpc.Name, "update-test-vpc") != 0 {
		t.Error("updated vpc does not have correct name")
	}

	if strings.Compare(uvpc.CidrBlock, "10.10.0.0/26") != 0 {
		t.Error("updated vpc does not have correct cidr block")
	}

	if uvpc.ID != vpc.ID {
		t.Error("updated vpc does not represent the correct db record")
	}

	err = conn.Close()
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteVPC(t *testing.T) {
	vpc := s.VPC{
		Name:         "delete-test-vpc",
		Access:       "public",
		Location:     "us-east4",
		Provider:     "GCP",
		Env:          "dev",
		CidrBlock:    "192.168.0.0/16",
		Type:         "vpc",
		NetworkMapID: 1,
	}

	conn, err := Conn()
	if err != nil {
		t.Error(err)
	}
	repo := New(conn)

	err = repo.CreateVPC(&vpc)
	if err != nil {
		t.Error(err)
	}
	err = repo.DeleteVPC(vpc.ID)
	if err != nil {
		t.Error(err)
	}

	var dvpc s.VPC
	err = conn.Where("id = ?", vpc.ID).First(&dvpc).Error
	if err == nil {
		t.Error("found a vpc that should have been deleted")
	}

	err = conn.Close()
	if err != nil {
		t.Error(err)
	}
}
