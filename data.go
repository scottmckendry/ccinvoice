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

type Service struct {
	ID       int
	Service  string
	Quantity int
	Price    float64
}

type Dog struct {
	ID        int
	Name      string
	OwnerName string
	Address   string
	City      string
	Email     string
	Grouping  int
	Services  []Service
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
            grouping INTEGER
        );
    `)
	if err != nil {
		return fmt.Errorf("error creating dogs table: %v", err)
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS dog_services (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            dog_id INTEGER,
            service TEXT,
            quantity INTEGER,
            price REAL,
            FOREIGN KEY(dog_id) REFERENCES dogs(id) ON DELETE CASCADE
        );
    `)
	if err != nil {
		return fmt.Errorf("error creating dog_services table: %v", err)
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
			&dog.Grouping,
		)
		if err != nil {
			rows.Close()
			return nil, fmt.Errorf("error scanning dog: %v", err)
		}
		dogs = append(dogs, dog)
	}

	// Get services
	for i, dog := range dogs {
		rows, err := db.Query(`
            SELECT id, service, quantity, price
            FROM dog_services
            WHERE dog_id = ?
        `, dog.ID)
		if err != nil {
			return nil, fmt.Errorf("error getting services: %v", err)
		}
		defer rows.Close()

		for rows.Next() {
			var svc Service
			err := rows.Scan(&svc.ID, &svc.Service, &svc.Quantity, &svc.Price)
			if err != nil {
				return nil, fmt.Errorf("error scanning service: %v", err)
			}
			dogs[i].Services = append(dogs[i].Services, svc)
		}
	}

	return dogs, nil
}

func getDog(id int) (Dog, error) {
	var dog Dog
	err := db.QueryRow(`
        SELECT id, name, ownerName, address, city, email, grouping 
        FROM dogs WHERE id = ?`, id).Scan(
		&dog.ID,
		&dog.Name,
		&dog.OwnerName,
		&dog.Address,
		&dog.City,
		&dog.Email,
		&dog.Grouping,
	)
	if err != nil {
		return Dog{}, fmt.Errorf("error getting dog: %v", err)
	}

	// Get services
	rows, err := db.Query(`
        SELECT id, service, quantity, price 
        FROM dog_services 
        WHERE dog_id = ?`, id)
	if err != nil {
		return Dog{}, fmt.Errorf("error getting services: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var svc Service
		err := rows.Scan(&svc.ID, &svc.Service, &svc.Quantity, &svc.Price)
		if err != nil {
			return Dog{}, fmt.Errorf("error scanning service: %v", err)
		}
		dog.Services = append(dog.Services, svc)
	}

	return dog, nil
}

func addDog(dog Dog) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("error beginning transaction: %v", err)
	}

	result, err := tx.Exec(`
        INSERT INTO dogs (name, ownerName, address, city, email, grouping)
        VALUES (?, ?, ?, ?, ?, ?)
    `, dog.Name, dog.OwnerName, dog.Address, dog.City, dog.Email, dog.Grouping)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error adding dog: %v", err)
	}

	dogID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error getting last insert id: %v", err)
	}

	for _, svc := range dog.Services {
		_, err = tx.Exec(`
            INSERT INTO dog_services (dog_id, service, quantity, price)
            VALUES (?, ?, ?, ?)
        `, dogID, svc.Service, svc.Quantity, svc.Price)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error adding service: %v", err)
		}
	}

	return tx.Commit()
}

func updateDog(dog Dog) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("error beginning transaction: %v", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Update dog information
	_, err = tx.Exec(`
        UPDATE dogs SET
            name = ?,
            ownerName = ?,
            address = ?,
            city = ?,
            email = ?,
            grouping = ?
        WHERE id = ?
    `, dog.Name, dog.OwnerName, dog.Address, dog.City, dog.Email, dog.Grouping, dog.ID)
	if err != nil {
		return fmt.Errorf("error updating dog: %v", err)
	}

	// Delete existing services for this dog
	_, err = tx.Exec(`DELETE FROM dog_services WHERE dog_id = ?`, dog.ID)
	if err != nil {
		return fmt.Errorf("error deleting existing services: %v", err)
	}

	// Insert new services
	for _, service := range dog.Services {
		_, err = tx.Exec(`
            INSERT INTO dog_services (dog_id, service, quantity, price)
            VALUES (?, ?, ?, ?)
        `, dog.ID, service.Service, service.Quantity, service.Price)
		if err != nil {
			return fmt.Errorf("error inserting service: %v", err)
		}
	}

	return tx.Commit()
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

func markEmailsInProcess(emails []Email) error {
	errs := []string{}
	for _, email := range emails {
		_, err := db.Exec("UPDATE email_queue SET sent = datetime('now') WHERE id = ?", email.ID)
		if err != nil {
			errs = append(errs, fmt.Sprintf("error marking email in process: %v", err))
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("%s", strings.Join(errs, "\n"))
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
