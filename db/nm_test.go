package db

import (
	"strings"
	"testing"

	s "github.com/thedevelopnik/netplan/structs"
)

func TestCreateAndDeleteNetworkMap(t *testing.T) {
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
		t.Error(err)
	}

	err = conn.Close()
	if err != nil {
		t.Error(err)
	}
}