package postgres

import (
	internal_types "fund-management-information-system/internal_types"
	"fund-management-information-system/pkg/repository/utils"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

type Account struct {
	Id       int    `db:"id"`
	Login    string `db:"login"`
	Password string `db:"password_hash"`
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateClient(client internal_types.SignUpClient) (int, error) {
	var accountId, personId, investmentId, clientId int
	addAccountQuery := `
	INSERT INTO accounts
		(login, password_hash) 
	VALUES
		($1, $2)
	RETURNING id
`
	addPersonQuery := `
	INSERT INTO persons
		(email, phone, address) 
	VALUES
		($1, '', '')
	RETURNING id;
`
	addInvestmentsInfoQuery := `
	INSERT INTO client_investments_info
		(bank_account, investment_amount, investment_strategy) 
	VALUES
		('', 0.0, '') 
	RETURNING id
`
	addClientQuery := `
	INSERT INTO clients
		(manager_id, account_id, person_id, client_investments_info_id) 
	VALUES 
		($1, $2, $3, $4)
	RETURNING id
`

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	if err := utils.EmailBusy(client.Email, r.db); err != nil {
		return 0, err
	}
	if err := utils.LoginBusy(client.Login, r.db); err != nil {
		return 0, err
	}

	err = r.db.QueryRow(addAccountQuery, client.Login, client.Password).Scan(&accountId)
	if err != nil {
		return 0, tx.Rollback()
	}

	err = r.db.QueryRow(addPersonQuery, client.Email).Scan(&personId)
	if err != nil {
		return 0, tx.Rollback()
	}

	err = r.db.QueryRow(addInvestmentsInfoQuery).Scan(&investmentId)
	if err != nil {
		return 0, tx.Rollback()
	}

	err = r.db.QueryRow(addClientQuery, client.ManagerId, accountId, personId, investmentId).Scan(&clientId)
	if err != nil {
		return 0, tx.Rollback()
	}

	return clientId, tx.Commit()

}

func (r *AuthPostgres) CreateManagerAccount(manager internal_types.ManagerAccount) (int, error) {
	var accountId, personId, workInfoId, managerId int

	addAccountQuery := `
	INSERT INTO accounts
	    (login, password_hash) 
	VALUES 
	    ($1, $2)
	RETURNING id
	`
	addPersonQuery := `	
	INSERT INTO persons
	    (email)
	VALUES 
	    ($1)
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
		return 0, tx.Rollback()
	}

	err = r.db.QueryRow(addPersonQuery, manager.Email).Scan(&personId)
	if err != nil {
		return 0, tx.Rollback()
	}

	err = r.db.QueryRow(addWorkInfoQuery).Scan(&workInfoId)
	if err != nil {
		return 0, tx.Rollback()
	}
	err = r.db.QueryRow(addManagerQuery, accountId, personId, workInfoId).Scan(&managerId)
	if err != nil {
		return 0, tx.Rollback()
	}
	if managerId == 0 {
		if err != nil {
			return 0, tx.Rollback()
		}
	}
	return managerId, tx.Commit()
}

func (r *AuthPostgres) GetAccount(login, password string) (Account, error) {
	var user Account

	query := `
	SELECT id, login, password_hash
	FROM accounts
	WHERE accounts.login=$1 AND accounts.password_hash=$2
`

	err := r.db.Get(&user, query, login, password)
	return user, err
}
