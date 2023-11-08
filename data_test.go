package main

import (
	"testing"
)

var testDog Dog = Dog{
	Name:      "Fido",
	OwnerName: "John Doe",
	Address:   "123 Fake Street",
	City:      "Fakeville",
	Email:     "johndoe@example.com",
	Service:   "walk",
	Quantity:  1,
	Price:     25,
}

func TestInit(t *testing.T) {
	err := Init()
	if err != nil {
		t.Errorf("Init() error = %q", err)
	}

	if Db == nil {
		t.Errorf("Init() Db is nil")
	}
}

func TestConnect(t *testing.T) {
	err := connect()
	if err != nil {
		t.Errorf("connect() error = %q", err)
	}

	if Db == nil {
		t.Errorf("connect() Db is nil")
	}
}

func TestCreateTables(t *testing.T) {
	err := createTables()
	if err != nil {
		t.Errorf("createTables() error = %q", err)
	}

	_, err = Db.Exec(`
        INSERT INTO dogs (
            name,
            ownerName,
            address,
            city,
            email,
            service,
            quantity,
            price
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
    `, testDog.Name, testDog.OwnerName, testDog.Address, testDog.City, testDog.Email, testDog.Service, testDog.Quantity, testDog.Price)

	if err != nil {
		t.Errorf("createTables() error = %q", err)
	}
}

func TestAddDog(t *testing.T) {
	err := AddDog(testDog)
	if err != nil {
		t.Errorf("AddDog() error = %q", err)
	}
}

func TestUpdateDog(t *testing.T) {
	testDog.ID = 1
	testDog.Name = "Fred"
	err := UpdateDog(testDog)
	if err != nil {
		t.Errorf("UpdateDog() error = %q", err)
	}

	dog, err := GetDog(1)
	if err != nil {
		t.Errorf("GetDog() error = %q", err)
	}

	if dog.Name != "Fred" {
		t.Errorf("UpdateDog() failed to update dog")
	}
}

func TestGetDogs(t *testing.T) {
	dogs, err := GetDogs()
	if err != nil {
		t.Errorf("GetDogs() error = %q", err)
	}

	if len(dogs) < 1 {
		t.Errorf("GetDogs() returned no dogs")
	}
}

func TestGetDog(t *testing.T) {
	dog, err := GetDog(1)
	if err != nil {
		t.Errorf("GetDog() error = %q", err)
	}

	if dog.ID != 1 {
		t.Errorf("GetDog() returned wrong dog")
	}
}

func TestDeleteDog(t *testing.T) {
	err := DeleteDog(1)
	if err != nil {
		t.Errorf("DeleteDog() error = %q", err)
	}

	_, err = GetDog(1)
	if err == nil {
		t.Errorf("DeleteDog() failed to delete dog")
	}

	t.Cleanup(func() {
		Db.Exec("DROP TABLE dogs")
	})
}
