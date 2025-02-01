package main

import (
	"os"
	"testing"
)

var testDog = Dog{
	ID:        1,
	Name:      "Fido",
	OwnerName: "John Doe",
	Address:   "123 Fake Street",
	City:      "Fakeville",
	Email:     "noreply@scottmckendry.tech",
	Grouping:  2,
	Services:  []Service{},
}

func dogsEqual(a, b Dog) bool {
	if a.ID != b.ID || a.Name != b.Name || a.OwnerName != b.OwnerName ||
		a.Address != b.Address || a.City != b.City || a.Email != b.Email ||
		a.Grouping != b.Grouping || len(a.Services) != len(b.Services) {
		return false
	}

	for i := range a.Services {
		if a.Services[i].Service != b.Services[i].Service ||
			a.Services[i].Quantity != b.Services[i].Quantity ||
			a.Services[i].Price != b.Services[i].Price {
			return false
		}
	}
	return true
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

	// Insert test dog
	_, err = db.Exec(`
        INSERT INTO dogs (
            name,
            ownerName,
            address,
            city,
            email,
            grouping
        ) VALUES (?, ?, ?, ?, ?, ?)
    `, testDog.Name, testDog.OwnerName, testDog.Address, testDog.City, testDog.Email, testDog.Grouping)

	if err != nil {
		t.Errorf("error inserting test dog: %q", err)
	}

	// Insert test service
	_, err = db.Exec(`
        INSERT INTO dog_services (
            dog_id,
            service,
            quantity,
            price
        ) VALUES (?, ?, ?, ?)
    `, testDog.ID, "Bath", 1, 20.00)

	if err != nil {
		t.Errorf("error inserting test service: %q", err)
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
		db.Exec("DROP TABLE dog_services")
		db.Exec("DROP TABLE dogs")
		Init()
	})
}

func TestUpdateTables(t *testing.T) {
	err := updateTables()
	if err != nil {
		t.Errorf("updateTables() error = %q", err)
	}

	// Force error
	oldDbUrl := dbUrl
	dbUrl = "someBadUrl"
	_ = connect()
	err = updateTables()
	if err == nil {
		t.Errorf("no error returned when expected")
	}

	t.Cleanup(func() {
		dbUrl = oldDbUrl
		_ = connect()
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

	if !dogsEqual(dog, testDog) {
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
	updatedDog := testDog
	updatedDog.Name = "Fred"
	err := updateDog(updatedDog)
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
	err = updateDog(updatedDog)

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

	if !dogsEqual(dogs[0], testDog) {
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
	db.Exec("DROP TABLE dog_services")
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

	if !dogsEqual(dog, testDog) {
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

func TestLoadEnv(t *testing.T) {
	// Create .env file if it doesn't exist
	_, err := os.Stat(".env")
	if os.IsNotExist(err) {
		_, err := os.Create(".env")
		if err != nil {
			t.Errorf("error creating .env file: %v", err)
		}
	}

	err = loadEnv()
	if err != nil {
		t.Errorf("loadEnv() error = %q", err)
	}

	// Rename .env file to force error
	err = os.Rename(".env", ".env.bak")
	if err != nil {
		t.Errorf("error renaming .env file: %v", err)
	}

	err = loadEnv()
	if err == nil {
		t.Errorf("no error returned when expected")
	}

	t.Cleanup(func() {
		err = os.Rename(".env.bak", ".env")
		if err != nil {
			t.Errorf("error renaming .env file: %v", err)
		}
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
            grouping integer
        );
    `)
	if err != nil {
		return err
	}

	return nil
}
