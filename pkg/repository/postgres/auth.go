package postgres

import (
	"errors"
	"fmt"
	internal_types "fund-management-information-system/internal-types"
	"github.com/jmoiron/sqlx"
)

const (
	clientsTable  = "clients"
	managersTable = "managers"
)
const (
	loginExistErr = "Логин уже занят"
	phoneExistErr = "Номер телефона уже занят"
)

type AuthPostgres struct {
	db *sqlx.DB
}
type User struct {
	Id       int
	Login    string
	Password string
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateClient(client internal_types.Client) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO clients (Name, Surname, Address, Phone, Email, Login, PasswordHash) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING Id")

	row := r.db.QueryRow(query, client.Name, client.Surname, client.Address, client.Phone, client.Email, client.Login, client.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) CreateManager(manager internal_types.Manager) (int, error) {
	if loginExist(managersTable, manager.Login, r.db) {
		return 0, errors.New(loginExistErr)
	}
	if phoneExist(managersTable, manager.Phone, r.db) {
		return 0, errors.New(phoneExistErr)
	}

	var id int
	query := fmt.Sprintf("INSERT INTO %s (Name, Surname, Address, Email, Phone, Login, PasswordHash) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING Id", managersTable)
	row := r.db.QueryRow(query, manager.Name, manager.Surname, manager.Address, manager.Email, manager.Phone, manager.Login, manager.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) Manager(login, password string) (internal_types.Manager, error) {
	var manager internal_types.Manager

	query := fmt.Sprintf("SELECT Id FROM managers WHERE Login=$1 AND PasswordHash=$2")
	err := r.db.Get(&manager, query, login, password)

	return manager, err
}

func (r *AuthPostgres) Client(login, password string) (internal_types.Client, error) {
	var client internal_types.Client

	query := fmt.Sprintf("SELECT Id FROM clients WHERE Login=$1 AND PasswordHash=$2")
	err := r.db.Get(&client, query, login, password)

	return client, err
}

func (r *AuthPostgres) User(login, password string) (User, error) {
	var user User

	query := fmt.Sprintf("SELECT Id FROM managers WHERE login=$1 AND passwordhash=$2 UNION SELECT Id FROM clients WHERE login=$1 AND passwordhash=$2")
	err := r.db.Get(&user, query, login, password)

	return user, err
}

func loginExist(table, login string, db *sqlx.DB) bool {
	var exist bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE login=$1)", table)
	row := db.QueryRow(query, login)
	_ = row.Scan(&exist)
	return exist
}
func phoneExist(table, phone string, db *sqlx.DB) bool {
	var exist bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE phone=$1)", table)
	row := db.QueryRow(query, phone)
	_ = row.Scan(&exist)
	return exist
}
