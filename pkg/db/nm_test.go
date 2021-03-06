package db

import (
	"reflect"
	"strings"
	"testing"

	s "github.com/thedevelopnik/netplan/pkg/models"
)

func TestCreateNetworkMap(t *testing.T) {
	nm := s.NetworkMap{
		Name: "create-test-nm",
	}

	conn, err := Conn()
	if err != nil {
		t.Error(err)
	}
	repo := New(conn)

	err = repo.CreateNetworkMap(&nm)
	if err != nil {
		t.Error(err)
	}

	if nm.ID == 0 {
		t.Error("network map did not have valid ID")
	}
	if strings.Compare(nm.Name, "create-test-nm") != 0 {
		t.Error("network map name was not correct")
	}

	err = conn.Close()
	if err != nil {
		t.Error(err)
	}
}

func TestGetNetworkMap(t *testing.T) {
	nm := s.NetworkMap{
		Name: "get-test-nm",
	}

	conn, err := Conn()
	if err != nil {
		t.Error(err)
	}
	repo := New(conn)

	err = repo.CreateNetworkMap(&nm)
	if err != nil {
		t.Error(err)
	}

	gnm, err := repo.GetNetworkMap(nm.ID)
	if err != nil {
		t.Error(err)
	}
	if gnm.ID != nm.ID {
		t.Error("retrieved id did not match original id")
	}

	if strings.Compare(gnm.Name, nm.Name) != 0 {
		t.Error("retrieved name did not match original name")
	}

	err = conn.Close()
	if err != nil {
		t.Error(err)
	}
}

func TestGetAllNetworkMaps(t *testing.T) {
	nm1 := s.NetworkMap{
		Name: "get-test-nm-1",
	}
	nm2 := s.NetworkMap{
		Name: "get-test-nm-2",
	}

	conn, err := Conn()
	if err != nil {
		t.Error(err)
	}
	repo := New(conn)

	err = repo.CreateNetworkMap(&nm1)
	if err != nil {
		t.Error(err)
	}

	err = repo.CreateNetworkMap(&nm2)
	if err != nil {
		t.Error(err)
	}

	networkMaps, err := repo.GetAllNetworkMaps()
	if err != nil {
		t.Error(err)
	}

	if len(networkMaps) < 2 {
		t.Errorf("expected networkmaps length to be at least 2, instead it was %d", len(networkMaps))
	}

	if reflect.TypeOf(networkMaps) != reflect.TypeOf([]s.NetworkMap{}) {
		t.Errorf("expected a slice of networkMaps by value, got %v", reflect.TypeOf(networkMaps))
	}
}

func TestUpdateNetworkMap(t *testing.T) {
	nm := s.NetworkMap{
		Name: "original-test-nm",
	}

	conn, err := Conn()
	if err != nil {
		t.Error(err)
	}
	repo := New(conn)

	err = repo.CreateNetworkMap(&nm)
	if err != nil {
		t.Error(err)
	}
	updater := s.NetworkMap{}
	updater.ID = nm.ID
	updater.Name = "update-test-nm"

	unm, err := repo.UpdateNetworkMap(&updater)

	if strings.Compare(unm.Name, "update-test-nm") != 0 {
		t.Error("updated networkmap does not have correct name")
	}

	if unm.ID != nm.ID {
		t.Error("updated network map does not represent the correct db record")
	}

	err = conn.Close()
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteNetworkMap(t *testing.T) {
	nm := s.NetworkMap{
		Name: "delete-test-nm",
	}

	conn, err := Conn()
	if err != nil {
		t.Error(err)
	}
	repo := New(conn)

	err = repo.CreateNetworkMap(&nm)
	if err != nil {
		t.Error(err)
	}
	err = repo.DeleteNetworkMap(nm.ID)
	if err != nil {
		t.Error(err)
	}

	_, err = repo.GetNetworkMap(nm.ID)
	if err == nil {
		t.Error("found a network map that should have been deleted")
	}

	err = conn.Close()
	if err != nil {
		t.Error(err)
	}
}

func TestGetNetworkMapWithVPCAndSubnets(t *testing.T) {
	nm := s.NetworkMap{
		Name: "full-test-nm",
	}

	conn, err := Conn()
	if err != nil {
		t.Error(err)
	}
	repo := New(conn)

	err = repo.CreateNetworkMap(&nm)
	if err != nil {
		t.Error(err)
	}

	vpc := s.VPC{
		Name:         "full-test-vpc",
		Access:       "public",
		Location:     "us-east4",
		Provider:     "GCP",
		Env:          "dev",
		CidrBlock:    "192.168.0.0/16",
		Type:         "vpc",
		NetworkMapID: nm.ID,
	}

	err = repo.CreateVPC(&vpc)
	if err != nil {
		t.Error(err)
	}

	sn := s.Subnet{
		Name:      "full-test-subnet",
		Access:    "public",
		Location:  "us-east4",
		Provider:  "GCP",
		Env:       "dev",
		CidrBlock: "192.168.0.0/24",
		VPCID:     vpc.ID,
	}

	err = repo.CreateSubnet(&sn)
	if err != nil {
		t.Error(err)
	}

	retrievedNM, err := repo.GetNetworkMap(nm.ID)

	if len(retrievedNM.VPCs) == 0 {
		t.Error("network map has no vpcs")
		return
	}

	retrievedVPC := retrievedNM.VPCs[0]
	if strings.Compare(retrievedVPC.Name, "full-test-vpc") != 0 {
		t.Error("vpc belonging to network map is not the correct vpc")
	}

	if len(retrievedVPC.Subnets) == 0 {
		t.Error("vpc has no subnets")
		return
	}

	retrievedSubnet := retrievedVPC.Subnets[0]
	if strings.Compare(retrievedSubnet.Name, "full-test-subnet") != 0 {
		t.Error("subnet belonging to vpc is not the correct subnet")
	}

	err = conn.Close()
	if err != nil {
		t.Error(err)
	}
}
