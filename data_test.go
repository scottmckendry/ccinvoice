package main

import (
	"testing"
)

var testDog Dog = Dog{
	ID:        1,
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

	if db == nil {
		t.Errorf("Init() Db is nil")
	}

	// Force error
	oldDbUrl := dbUrl
	dbUrl = "someBadUrl"
	err = Init()
	if err == nil {
		t.Errorf("no error returned when expected")
	}

	t.Cleanup(func() {
		dbUrl = oldDbUrl
		Init()
	})
}

func TestConnect(t *testing.T) {
	err := connect()
	if err != nil {
		t.Errorf("connect() error = %q", err)
	}

	if db == nil {
		t.Errorf("connect() Db is nil")
	}

	// Force error
	oldDbUrl := dbUrl
	dbUrl = "someBadUrl"
	err = connect()
	if err == nil {
		t.Errorf("no error returned when expected")
	}

	t.Cleanup(func() {
		dbUrl = oldDbUrl
		Init()
	})
}

func TestCreateTables(t *testing.T) {
	err := createTables()
	if err != nil {
		t.Errorf("createTables() error = %q", err)
	}

	_, err = db.Exec(`
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

	// Force error
	oldDbUrl := dbUrl
	dbUrl = "someBadUrl"
	_ = connect()
	err = createTables()
	if err == nil {
		t.Errorf("no error returned when expected")
	}

	t.Cleanup(func() {
		dbUrl = oldDbUrl
		_ = connect()
		db.Exec("DROP TABLE dogs")
		Init()
	})
}

func TestAddDog(t *testing.T) {
	err := addDog(testDog)
	if err != nil {
		t.Errorf("AddDog() error = %q", err)
	}

	dog, err := getDog(1)
	if err != nil {
		t.Errorf("GetDog() error = %q", err)
	}

	if dog != testDog {
		t.Errorf("AddDog() failed to add dog")
		t.Errorf("have: %v", dog)
		t.Errorf("want: %v", testDog)
	}

	// Add bad data
	badDog := Dog{
		Name: "Fido", // Column is unique
	}
	err = addDog(badDog)
	if err == nil {
		t.Errorf("no error returned when expected")
	}
}

func TestUpdateDog(t *testing.T) {
	testDog.ID = 1
	testDog.Name = "Fred"
	err := updateDog(testDog)
	if err != nil {
		t.Errorf("UpdateDog() error = %q", err)
	}

	dog, err := getDog(1)
	if err != nil {
		t.Errorf("GetDog() error = %q", err)
	}

	if dog.Name != "Fred" {
		t.Errorf("UpdateDog() failed to update dog")
	}

	// Force error
	db.Exec("DROP TABLE dogs")
	err = updateDog(testDog)

	if err == nil {
		t.Errorf("no error returned when expected")
	}

	t.Cleanup(func() {
		Init()
		_ = addDog(testDog)
	})
}

func TestGetDogs(t *testing.T) {
	dogs, err := getDogs()
	if err != nil {
		t.Errorf("GetDogs() error = %q", err)
	}

	if len(dogs) < 1 {
		t.Errorf("GetDogs() returned no dogs")
	}

	if dogs[0] != testDog {
		t.Errorf("GetDogs() returned incorrect dog")
		t.Errorf("have: %v", dogs[0])
		t.Errorf("want: %v", testDog)
	}

	// Force scan error
	db.Exec("insert into dogs (name) values (NULL)")
	_, err = getDogs()
	if err == nil {
		t.Errorf("no error returned when expected")
	}

	// Force query error
	_, err = db.Exec("drop table dogs;")
	if err != nil {
		t.Errorf("error dropping table: %v", err)
	}
	_, err = getDogs()
	if err == nil {
		t.Errorf("no error returned when expected")
	}

	t.Cleanup(func() {
		Init()
		_ = addDog(testDog)
	})

}

func TestGetDog(t *testing.T) {
	dog, err := getDog(1)
	if err != nil {
		t.Errorf("GetDog() error = %q", err)
	}

	if dog != testDog {
		t.Errorf("GetDog() returned incorrect dog")
	}
}

func TestDeleteDog(t *testing.T) {
	err := deleteDog(1)
	if err != nil {
		t.Errorf("DeleteDog() error = %q", err)
	}

	_, err = getDog(1)
	if err == nil {
		t.Errorf("DeleteDog() failed to delete dog")
	}

	// Force error
	db.Exec("DROP TABLE dogs")
	err = deleteDog(1)
	if err == nil {
		t.Errorf("no error returned when expected")
	}

	t.Cleanup(func() {
		Init()
		_ = addDog(testDog)
	})
}

func createBadTable() error {
	_, err := db.Exec(`
        CREATE TABLE dogs (
            id integer PRIMARY KEY,
            name text,
            ownerName text,
            address text,
            city text,
            email text,
            service text,
            quantity text,
            price text
        );
    `)
	if err != nil {
		return err
	}

	return nil
}
