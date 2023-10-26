package postgres

import (
	"fund-management-information-system/internal_types"
	"github.com/jmoiron/sqlx"
)

type ClientRepositoryPostgres struct {
	db *sqlx.DB
}

func NewClientPostgres(db *sqlx.DB) *ClientRepositoryPostgres {
	return &ClientRepositoryPostgres{db: db}
}
func (r *ClientRepositoryPostgres) UpdateClient(id int, client internal_types.Client) error {
	var err error

	updateClientQuery := `
	UPDATE persons
	SET email=$1, phone=$2, address = $3
	WHERE id IN (SELECT person_id FROM clients WHERE id=$4)
`
	updateInvestmentsQuery := `
	UPDATE client_investments_info
	SET bank_account=$1, investment_amount=$2, investment_strategy = $3
	WHERE id IN (SELECT client_investments_info FROM clients WHERE clients.id=$4)
`
	_, err = r.db.Exec(updateClientQuery, client.Email, client.Phone, client.Address, id)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(updateInvestmentsQuery, client.BankAccount, client.InvestmentAmount, client.InvestmentStrategy, id)
	if err != nil {
		return err
	}

	return nil
}
func (r *ClientRepositoryPostgres) UpdateInvestments(client internal_types.Client) error {
	var err error

	updateCIIQuery := `
	UPDATE client_investments_info
	SET investment_amount=$1
	WHERE id = (
	    SELECT id
	    FROM client_investments_info
	    ORDER BY random()
	    LIMIT 1
	)
`
	_, err = r.db.Exec(updateCIIQuery, client.InvestmentAmount)
	if err != nil {
	}

	return nil
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
func (r *ClientRepositoryPostgres) GetById(clientId int) (internal_types.Client, error) {
	var client internal_types.Client

	query := `
	SELECT client.id, person.email, person.phone, person.email,
	       invest.bank_account, invest.investment_amount, invest.investment_strategy
	FROM clients client
	LEFT JOIN persons person ON client.id = person.id
	LEFT JOIN client_investments_info invest ON client.id = invest.id
	WHERE client.id=$1
`
	if err := r.db.Get(&client, query, clientId); err != nil {
		return client, err
	}

	return client, nil
}
