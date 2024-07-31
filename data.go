package main

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
	_ "modernc.org/sqlite"
)

var dbUrl = "file:./db.sqlite3"
var db *sql.DB

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
	Grouping  int
}

type Email struct {
	ID     int
	DogID  int
	Queued sql.NullString
	Sent   sql.NullString
}

func Init() error {
	err := connect()
	if err != nil {
		return err
	}

	err = createTables()
	if err != nil {
		return fmt.Errorf("error creating tables: %v", err)
	}

	err = updateTables()
	if err != nil {
		return fmt.Errorf("error updating tables: %v", err)
	}

	return nil
}

func connect() error {
	db, _ = sql.Open("libsql", dbUrl)

	err := db.Ping()
	if err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}

	return nil
}

func createTables() error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS dogs (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT unique,
            ownerName TEXT,
            address TEXT,
            city TEXT,
            email TEXT,
            service TEXT,
            quantity INTEGER,
            price INTEGER
            grouping INTEGER
        );
    `)
	if err != nil {
		return fmt.Errorf("error creating dogs table: %v", err)
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS email_queue (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            dog_id INTEGER,
            queued TEXT,
            sent TEXT
        );
    `)

	if err != nil {
		return fmt.Errorf("error creating email_queue table: %v", err)
	}

	return nil
}

func updateTables() error {
	_, err := db.Exec(`
        ALTER TABLE dogs ADD COLUMN grouping INTEGER;
    `)
	if err != nil && !strings.Contains(err.Error(), "duplicate column name") {
		return fmt.Errorf("error updating dogs table: %v", err)
	}

	_, err = db.Exec(`
        UPDATE dogs SET grouping = 0 WHERE grouping IS NULL
    `)
	if err != nil {
		return fmt.Errorf("error updating dogs table: %v", err)
	}

	return nil
}

func getDogs() ([]Dog, error) {
	rows, err := db.Query("SELECT * FROM dogs order by grouping")
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
			&dog.Grouping,
		)
		if err != nil {
			rows.Close()
			return nil, fmt.Errorf("error scanning dog: %v", err)
		}
		dogs = append(dogs, dog)
	}

	return dogs, nil
}

func getDog(id int) (Dog, error) {
	var dog Dog
	err := db.QueryRow("SELECT * FROM dogs WHERE id = ?", id).Scan(
		&dog.ID,
		&dog.Name,
		&dog.OwnerName,
		&dog.Address,
		&dog.City,
		&dog.Email,
		&dog.Service,
		&dog.Quantity,
		&dog.Price,
		&dog.Grouping,
	)
	if err != nil {
		return Dog{}, fmt.Errorf("error getting dog: %v", err)
	}

	return dog, nil
}

func addDog(dog Dog) error {
	_, err := db.Exec(`
        INSERT INTO dogs (
            name,
            ownerName,
            address,
            city,
            email,
            service,
            quantity,
            price,
            grouping
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
    `, dog.Name, dog.OwnerName, dog.Address, dog.City, dog.Email, dog.Service, dog.Quantity, dog.Price, dog.Grouping)
	if err != nil {
		return fmt.Errorf("error adding dog: %v", err)
	}

	return nil
}

func updateDog(dog Dog) error {
	_, err := db.Exec(`
        UPDATE dogs SET
            name = ?,
            ownerName = ?,
            address = ?,
            city = ?,
            email = ?,
            service = ?,
            quantity = ?,
            price = ?,
            grouping = ?
        WHERE id = ?
    `, dog.Name, dog.OwnerName, dog.Address, dog.City, dog.Email, dog.Service, dog.Quantity, dog.Price, dog.Grouping, dog.ID)
	if err != nil {
		return fmt.Errorf("error updating dog: %v", err)
	}

	return nil
}

func deleteDog(id int) error {
	_, err := db.Exec("DELETE FROM dogs WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("error deleting dog: %v", err)
	}

	return nil
}

func queueEmail(dogID int) error {
	_, err := db.Exec(`
        INSERT INTO email_queue (
            dog_id, queued, sent
        ) VALUES (?, datetime('now'), NULL)
    `, dogID)
	if err != nil {
		return fmt.Errorf("error queuing email: %v", err)
	}

	return nil
}

func getEmailQueue() ([]Email, error) {
	rows, err := db.Query("SELECT * FROM email_queue WHERE sent IS NULL")
	if err != nil {
		return nil, fmt.Errorf("error getting email queue: %v", err)
	}

	var emails []Email
	for rows.Next() {
		var email Email
		err := rows.Scan(
			&email.ID,
			&email.DogID,
			&email.Queued,
			&email.Sent,
		)
		if err != nil {
			rows.Close()
			return nil, fmt.Errorf("error scanning email: %v", err)
		}
		emails = append(emails, email)
	}

	return emails, nil
}

func markEmailInProcess(id int) error {
	_, err := db.Exec("UPDATE email_queue SET sent = 'in_process' WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("error marking email in process: %v", err)
	}

	return nil
}

func markEmailSent(id int) error {
	_, err := db.Exec("UPDATE email_queue SET sent = datetime('now') WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("error marking email sent: %v", err)
	}

	return nil
}
