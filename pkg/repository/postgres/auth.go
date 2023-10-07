package postgres

import (
	"errors"
	"fmt"
	internal_types "fund-management-information-system/internal_types"
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
const (
	roleManager = "manager"
	roleClient  = "client"
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
	if loginExist(clientsTable, client.Login, r.db) {
		return 0, errors.New(loginExistErr)
	}
	if phoneExist(clientsTable, client.Phone, r.db) {
		return 0, errors.New(phoneExistErr)
	}

	var id int
	query := fmt.Sprintf("INSERT INTO %s (Name, Surname, Address, Phone, Email, Login, PasswordHash, ManagerId) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING Id", clientsTable)

	row := r.db.QueryRow(query, client.Name, client.Surname, client.Address, client.Phone, client.Email, client.Login, client.Password, client.ManagerId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) CreateManager(manager internal_types.SignUp) (int, error) {
	var accountId, personId, managerId int

	addAccountQuery := `
	INSERT INTO accounts
	    (login, password_hash) 
	VALUES 
	    ($1, $2)
	`
	addPersonEmailQuery := `	
	INSERT INTO persons
	    (email)
	VALUES 
	    ($1)
	`
	addManagerQuery := `
	INSERT INTO managers
	    (account_id, person_id)
	VALUES 
	    ($1, $2)
	RETURNING id
	`

	getAccountIdQuery := `
	SELECT id FROM accounts
	WHERE login=$1
	`
	getPersonIdQuery := `
	SELECT id FROM persons
	WHERE email=$1
	`

	row := r.db.QueryRow(addAccountQuery, manager.Login, manager.Password) // вносим в таблицу accounts
	row = r.db.QueryRow(getAccountIdQuery, manager.Login)                  // получаем id созданной записи accounts
	if err := row.Scan(&accountId); err != nil {
		return 0, err
	}

	row = r.db.QueryRow(addPersonEmailQuery, manager.Email)
	row = r.db.QueryRow(getPersonIdQuery, manager.Email)
	if err := row.Scan(&personId); err != nil {
		return 0, err
	}

	row = r.db.QueryRow(addManagerQuery, accountId, personId)
	if err := row.Scan(&managerId); err != nil {
		return 0, err
	}

	return managerId, nil
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

func (r *AuthPostgres) User(login, password string) (User, string, error) {
	var user User
	var role string

	if isManager(login, password, r.db) {
		role = roleManager
	} else if isClient(login, password, r.db) {
		role = roleClient
	}
	query := fmt.Sprintf("SELECT Id FROM managers WHERE login=$1 AND passwordhash=$2 UNION SELECT Id FROM clients WHERE login=$1 AND passwordhash=$2")
	err := r.db.Get(&user, query, login, password)

	return user, role, err
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
func isManager(login, password string, db *sqlx.DB) bool {
	var manager bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM managers WHERE login=$1 AND passwordhash=$2)")
	row := db.QueryRow(query, login, password)
	_ = row.Scan(&manager)
	return manager
}
func isClient(login, password string, db *sqlx.DB) bool {
	var client bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM clients WHERE login=$1 AND passwordhash=$2)")
	row := db.QueryRow(query, login, password)
	_ = row.Scan(&client)
	return client
}
