package models

import (
	"database/sql"
	"fmt"

	//_ "github.com/mattn/go-sqlite3"
	_ "modernc.org/sqlite"
)

type HospitalStore interface {
	Get() ([]Hospital, error)
	Create(h Hospital) (ID string, err error)
	Delete(id string) error
	Update(h Hospital) error
}

type Hospital struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	City       string `json:"city"`
	Pin        string `json:"pin"`
	Country    string `json:"country"`
	Created_at string `json:"createdAt"`
}

type HospitalSqlliteStore struct {
	db *sql.DB
}

func (p *HospitalSqlliteStore) Create(h Hospital) (ID string, err error) {
	// horrible SQL goes here
	// handle errors, etc
	res, err := p.db.Exec("INSERT INTO hospitals (name, address, city, country, pin) VALUES (?, ?, ?,?,?)", h.Name, h.Address, h.City, h.Country, h.Pin)
	if err != nil {
		return "nil", err
	}
	id, err1 := res.LastInsertId()
	if err1 != nil {
		return "nil", err1
	}

	return fmt.Sprint(id), nil
}

func (p *HospitalSqlliteStore) Delete(id string) error {
	// horrible SQL goes here
	// handle errors, etc
	res, err := p.db.Exec("DELETE FROM hospitals WHERE id = ?", id)
	if err != nil {
		return err
	}
	deleted, err := res.RowsAffected()
	if deleted < 1 {
		return fmt.Errorf("Hospital with ID: %v cannot be deleted because it doesn't exist", id)
	}
	return err
}
func (p *HospitalSqlliteStore) Update(h Hospital) error {
	_, err := p.db.Exec("UPDATE hospitals SET name = ?, address = ?, ciry = ? , country, , pin = ? = ? WHERE id = ?", h.Name, h.Address, h.City, h.Country, h.Pin, h.Id)
	return err
}

func (p *HospitalSqlliteStore) Get() ([]Hospital, error) {
	rows, err := p.db.Query("SELECT * FROM hospitals")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	hospitals := make([]Hospital, 0)
	for rows.Next() {
		aHospital := Hospital{}
		err = rows.Scan(&aHospital.Id, &aHospital.Name, &aHospital.Address, &aHospital.City, &aHospital.Country, &aHospital.Pin, &aHospital.Created_at)
		if err != nil {
			return nil, err
		}
		hospitals = append(hospitals, aHospital)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return hospitals, nil

}

func NewSqlliteStore() *HospitalSqlliteStore {

	db, err := sql.Open(envs["DB_DIALECT"], envs["DB"])
	if err != nil {
		panic(err)
	}
	return &HospitalSqlliteStore{
		db: db,
	}
}
