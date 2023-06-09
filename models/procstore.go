package models

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type MedProcs struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Desription string  `json:"description"`
	Cost       float32 `json:"cost"`
	Created_at string  `json:"createdAt"`
}

type ProcStore interface {
	Get() ([]MedProcs, error)
	Create(m MedProcs) (ID string, err error)
	Delete(id string) error
	Update(m MedProcs) error
}

type ProcSqlliteStore struct {
	DB *sql.DB
}

func (p *ProcSqlliteStore) Create(mp MedProcs) (ID string, err error) {
	// horrible SQL goes here
	// handle errors, etc
	res, err := p.DB.Exec("INSERT INTO medprocs (name, description, cost) VALUES (?, ?, ?)", mp.Name, mp.Desription, mp.Cost)

	if err != nil {
		return "nil", err
	}
	id, err1 := res.LastInsertId()
	if err1 != nil {
		return "nil", err1
	}

	return fmt.Sprint(id), nil
}

func (p *ProcSqlliteStore) Delete(id string) error {
	// horrible SQL goes here
	// handle errors, etc
	res, err := p.DB.Exec("DELETE FROM medprocs WHERE id = ?", id)
	if err != nil {
		return err
	}
	deleted, err := res.RowsAffected()
	if deleted < 1 {
		return fmt.Errorf("MedProc with ID: %v cannot be deleted because it doesn't exist", id)
	}
	return err
}
func (p *ProcSqlliteStore) Update(mp MedProcs) error {
	_, err := p.DB.Exec("UPDATE medprocs SET name = ?, description = ?, cost = ? WHERE id = ?", mp.Name, mp.Desription, mp.Cost, mp.Id)
	return err
}

func (p *ProcSqlliteStore) Get() ([]MedProcs, error) {
	rows, err := p.DB.Query("SELECT * FROM medprocs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	medProcs := make([]MedProcs, 0)
	for rows.Next() {
		amedProc := MedProcs{}
		err = rows.Scan(&amedProc.Id, &amedProc.Name, &amedProc.Desription, &amedProc.Cost, &amedProc.Created_at)
		if err != nil {
			return nil, err
		}
		medProcs = append(medProcs, amedProc)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return medProcs, nil

}
