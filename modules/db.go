package modules

import (
	"database/sql"
	"fmt"
)

type Database struct {
	connectionString string
}

func NewDatabase(connectionString string) *Database {
	return &Database{
		connectionString: connectionString,
	}
}

func (db *Database) Connect() (*sql.DB, error) {
	conn, err := sql.Open("postgres", db.connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	err = conn.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping the database: %v", err)
	}

	return conn, nil
}

func (db *Database) CreateTables() error {
	conn, err := db.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	// Create MPs table
	_, err = conn.Exec(`
		CREATE TABLE IF NOT EXISTS MPs (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			department TEXT NOT NULL
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create MPs table: %v", err)
	}

	// Create Local Counselors table
	_, err = conn.Exec(`
		CREATE TABLE IF NOT EXISTS LocalCounselors (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			region TEXT NOT NULL
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create LocalCounselors table: %v", err)
	}

	// Create Performance table
	_, err = conn.Exec(`
		CREATE TABLE IF NOT EXISTS Performance (
			id SERIAL PRIMARY KEY,
			mp_id INT REFERENCES MPs(id),
			local_counselor_id INT REFERENCES LocalCounselors(id),
			kpi TEXT NOT NULL,
			value INT NOT NULL,
			FOREIGN KEY (mp_id, kpi) REFERENCES MPs(id, kpi),
			FOREIGN KEY (local_counselor_id, kpi) REFERENCES LocalCounselors(id, kpi)
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create Performance table: %v", err)
	}

	return nil
}
