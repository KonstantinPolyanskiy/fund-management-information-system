package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ClientRepositoryPostgres struct {
	db *sqlx.DB
}

func NewClientPostgres(db *sqlx.DB) *ClientRepositoryPostgres {
	return &ClientRepositoryPostgres{db: db}
}

func (r *ClientRepositoryPostgres) Delete(clientId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE Id = $1", clientsTable)
	_, err := r.db.Exec(query, clientId)
	return err
}
