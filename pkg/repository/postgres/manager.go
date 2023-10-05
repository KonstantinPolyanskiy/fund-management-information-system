package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ManagerRepositoryPostgres struct {
	db *sqlx.DB
}

func NewManagerPostgres(db *sqlx.DB) *ManagerRepositoryPostgres {
	return &ManagerRepositoryPostgres{db: db}
}

func (r *ManagerRepositoryPostgres) Delete(managerId int) error {
	query := fmt.Sprintf("DELETE FROM managers WHERE id=$1")
	_, err := r.db.Exec(query, managerId)
	return err
}
