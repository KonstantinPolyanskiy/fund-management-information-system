package utils

import (
	"errors"
	"github.com/jmoiron/sqlx"
)

type Verifier interface {
	LoginBusy(login string, db *sqlx.DB) error
}

const (
	loginBusyErr = "Логин уже занят"
	phoneBusyErr = "Номер уже занят"
	emailBusyErr = "Почта уже занята"
)

func LoginBusy(login string, db *sqlx.DB) error {
	var busy bool
	query := `
	SELECT EXISTS 
	(
	SELECT 1 FROM accounts
	WHERE login=$1
	)
	`

	row := db.QueryRow(query, login)
	_ = row.Scan(&busy)
	if busy == true {
		return errors.New(loginBusyErr)
	}

	return nil
}

func EmailBusy(email string, db *sqlx.DB) error {
	var busy bool
	query := `
	SELECT EXISTS 
	(
	SELECT 1 FROM persons
	WHERE Email=$1 
	)
	`

	row := db.QueryRow(query, email)
	_ = row.Scan(&busy)
	if busy == true {
		return errors.New(emailBusyErr)
	}
	return nil
}
