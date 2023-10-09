package postgres

import (
	"errors"
	"fmt"
	internal_types "fund-management-information-system/internal_types"
	"fund-management-information-system/pkg/repository/utils"
	"github.com/jmoiron/sqlx"
)

const (
	accountsTable = "accounts"
	personsTable  = "persons"
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
	Id       int    `db:"id"`
	Login    string `db:"login"`
	Password string `db:"password_hash"`
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
	var accountId, personId, workInfoId, managerId int

	addAccountQuery := `
	INSERT INTO accounts
	    (login, password_hash) 
	VALUES 
	    ($1, $2)
	RETURNING id
	`
	addPersonEmailQuery := `	
	INSERT INTO persons
	    (email, phone, address)
	VALUES 
	    ($1, '', '')
	RETURNING id
	`
	addWorkInfoQuery := `
	INSERT INTO manager_work_info
		(bank_account, capital_managment, profit_percent_day) 
	VALUES
		('', 0, 0.0) 
	RETURNING id
`
	addManagerQuery := `
	INSERT INTO managers
	    (account_id, person_id, work_info_id)
	VALUES 
	    ($1, $2, $3)
	RETURNING id
	`

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	if err := utils.LoginBusy(manager.Login, r.db); err != nil {
		return 0, err
	}
	if err := utils.EmailBusy(manager.Email, r.db); err != nil {
		return 0, err
	}

	err = r.db.QueryRow(addAccountQuery, manager.Login, manager.Password).Scan(&accountId)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}

	err = r.db.QueryRow(addPersonEmailQuery, manager.Email).Scan(&personId)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}

	err = r.db.QueryRow(addWorkInfoQuery).Scan(&workInfoId)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}

	err = r.db.QueryRow(addManagerQuery, accountId, personId, workInfoId).Scan(&managerId)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}

	if managerId == 0 {
		err := tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, errors.New("Ошибка в создании пользователя")
	}
	return managerId, tx.Commit()
}
func (r *AuthPostgres) GetUser(login, password string) (User, error) {
	var user User

	query := `
	SELECT managers.id, login, password_hash
	FROM managers
	JOIN accounts ON managers.account_id = accounts.id
	WHERE accounts.login=$1 AND accounts.password_hash=$2
`

	err := r.db.Get(&user, query, login, password)
	return user, err
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
	_ = `
	SELECT m.id 	
	FROM managers m 
	JOIN accounts a ON a.id = m.account_id
	WHERE a.login=$1 AND a.password_hash=$2
`
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
