package postgres

import (
	"github.com/jmoiron/sqlx"
)

type ClientRepositoryPostgres struct {
	db *sqlx.DB
}

func NewClientPostgres(db *sqlx.DB) *ClientRepositoryPostgres {
	return &ClientRepositoryPostgres{db: db}
}

func (r *ClientRepositoryPostgres) DeleteById(clientId int) error {
	var managerId, accountId, InvestmentsId, personId int
	var err error

	getManagerIdQuery := `
	SELECT manager_id
	FROM clients
	JOIN managers ON clients.manager_id = managers.id
	WHERE clients.id=$1
`
	getAccountIdQuery := `
	SELECT account_id	 
	FROM clients
	JOIN accounts ON clients.account_id = accounts.id
	WHERE clients.id=$1
`
	getInvestmentsIdQuery := `
	SELECT client_investments_info_id
	FROM clients
	JOIN client_investments_info ON clients.client_investments_info_id = client_investments_info.id
	WHERE clients.id=$1
`
	getPersonIdQuery := `
	SELECT person_id
	FROM clients
	JOIN accounts ON clients.account_id = accounts.id
	WHERE clients.id=$1
`

	updateManagerIdQuery := `
	UPDATE clients 
	SET manager_id = NULL
	WHERE id=$1
`
	deleteClientQuery := `
	DELETE FROM clients
	WHERE id=$1
`
	deleteAccountIdQuery := `
	DELETE FROM accounts
	WHERE id=$1
`
	deleteInvestmentsIdQuery := `
	DELETE FROM client_investments_info
	WHERE id=$1
`
	deletePersonIdQuery := `
	DELETE FROM persons
	WHERE id=$1
`

	err = r.db.QueryRow(getManagerIdQuery, clientId).Scan(&managerId)
	if err != nil {
		return err
	}

	err = r.db.QueryRow(getAccountIdQuery, clientId).Scan(&accountId)
	if err != nil {
		return err
	}

	err = r.db.QueryRow(getInvestmentsIdQuery, clientId).Scan(&InvestmentsId)
	if err != nil {
		return err
	}

	err = r.db.QueryRow(getPersonIdQuery, clientId).Scan(&personId)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(deleteClientQuery, clientId)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(updateManagerIdQuery, managerId)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(deleteAccountIdQuery, accountId)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(deleteInvestmentsIdQuery, InvestmentsId)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(deletePersonIdQuery, personId)
	if err != nil {
		return err
	}

	return nil

}
