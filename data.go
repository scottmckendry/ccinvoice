package main

import (
	"database/sql"
	"fmt"

	_ "github.com/libsql/libsql-client-go/libsql"
	_ "modernc.org/sqlite"
)

var dbUrl = "file:./db.sqlite3"
var Db *sql.DB

type Dog struct {
	ID        int
	Name      string
	OwnerName string
	Address   string
	City      string
	Email     string
	Service   string
	Quantity  int
	Price     float64
}

func Init() error {
	err := connect()
	if err != nil {
		return err
	}

	err = createTables()
	if err != nil {
		return err
	}

	return nil
}

func connect() error {
	var err error
	Db, err = sql.Open("libsql", dbUrl)
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}

	return nil
}

func createTables() error {
	_, err := Db.Exec(`
        CREATE TABLE IF NOT EXISTS dogs (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT,
            ownerName TEXT,
            address TEXT,
            city TEXT,
            email TEXT,
            service TEXT,
            quantity INTEGER,
            price INTEGER
        );
    `)
	if err != nil {
		return fmt.Errorf("error creating dogs table: %v", err)
	}

	return nil
}

func GetDogs() ([]Dog, error) {
	rows, err := Db.Query("SELECT * FROM dogs")
	if err != nil {
		return nil, fmt.Errorf("error getting dogs: %v", err)
	}

	var dogs []Dog
	for rows.Next() {
		var dog Dog
		err := rows.Scan(
			&dog.ID,
			&dog.Name,
			&dog.OwnerName,
			&dog.Address,
			&dog.City,
			&dog.Email,
			&dog.Service,
			&dog.Quantity,
			&dog.Price,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning dog: %v", err)
		}
		dogs = append(dogs, dog)
	}

	return dogs, nil
}

func GetDog(id int) (Dog, error) {
	var dog Dog
	err := Db.QueryRow("SELECT * FROM dogs WHERE id = ?", id).Scan(
		&dog.ID,
		&dog.Name,
		&dog.OwnerName,
		&dog.Address,
		&dog.City,
		&dog.Email,
		&dog.Service,
		&dog.Quantity,
		&dog.Price,
	)
	if err != nil {
		return Dog{}, fmt.Errorf("error getting dog: %v", err)
	}

	return dog, nil
}

func AddDog(dog Dog) error {
	_, err := Db.Exec(`
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
    `, dog.Name, dog.OwnerName, dog.Address, dog.City, dog.Email, dog.Service, dog.Quantity, dog.Price)
	if err != nil {
		return fmt.Errorf("error adding dog: %v", err)
	}

	return nil
}

func UpdateDog(dog Dog) error {
	_, err := Db.Exec(`
        UPDATE dogs SET
            name = ?,
            ownerName = ?,
            address = ?,
            city = ?,
            email = ?,
            service = ?,
            quantity = ?,
            price = ?
        WHERE id = ?
    `, dog.Name, dog.OwnerName, dog.Address, dog.City, dog.Email, dog.Service, dog.Quantity, dog.Price, dog.ID)
	if err != nil {
		return fmt.Errorf("error updating dog: %v", err)
	}

	return nil
}

func DeleteDog(id int) error {
	_, err := Db.Exec("DELETE FROM dogs WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("error deleting dog: %v", err)
	}

	return nil
}
