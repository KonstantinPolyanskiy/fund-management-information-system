package postgres

import (
	"fund-management-information-system/internal_types"
	"github.com/jmoiron/sqlx"
)

type ManagerRepositoryPostgres struct {
	db *sqlx.DB
}

func NewManagerPostgres(db *sqlx.DB) *ManagerRepositoryPostgres {
	return &ManagerRepositoryPostgres{db: db}
}

func (r *ManagerRepositoryPostgres) DeleteById(managerId int) error {
	var accountId, workId, personId int
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	getAccountIdQuery := `
	SELECT account_id
	FROM managers
	JOIN accounts ON managers.account_id = accounts.id
	WHERE managers.id=$1
`
	getWorkIdQuery := `
	SELECT work_info_id
	FROM managers
	JOIN manager_work_info ON managers.work_info_id = manager_work_info.id
	WHERE managers.id=$1
`
	getPersonIdQuery := `
	SELECT person_id
	FROM managers 
	JOIN persons on managers.person_id = persons.id
	WHERE managers.id=$1
`

	deleteClientManagerQuery := `
	DELETE FROM clients
	WHERE manager_id=$1
`
	deleteWorkQuery := `
	DELETE FROM manager_work_info
	WHERE id=$1
`
	deletePersonQuery := `
	DELETE FROM persons
	WHERE id=$1
`
	deleteAccountQuery := `
	DELETE FROM accounts
	WHERE id=$1
`
	deleteManagerQuery := `
	DELETE FROM managers
	WHERE id=$1
`
	err = r.db.QueryRow(getAccountIdQuery, managerId).Scan(&accountId)
	err = r.db.QueryRow(getWorkIdQuery, managerId).Scan(&workId)
	err = r.db.QueryRow(getPersonIdQuery, managerId).Scan(&personId)

	_, err = r.db.Exec(deleteManagerQuery, managerId)
	_, err = r.db.Exec(deleteClientManagerQuery, managerId)
	_, err = r.db.Exec(deletePersonQuery, personId)
	_, err = r.db.Exec(deleteAccountQuery, accountId)
	_, err = r.db.Exec(deleteWorkQuery, workId)

	return tx.Commit()
}

func (r *ManagerRepositoryPostgres) GetById(managerId int) (internal_types.Manager, error) {
	var manager internal_types.Manager

	query := `
	SELECT mng.id, p.email, p.phone, p.address,
       mwi.bank_account, mwi.capital_managment, mwi.profit_percent_day
	FROM managers mng
	LEFT JOIN persons p ON mng.person_id = p.id
    LEFT JOIN manager_work_info mwi  ON mng.id = mwi.id
	WHERE mng.id=$1`

	if err := r.db.Get(&manager, query, managerId); err != nil {
		return manager, err
	}
	return manager, nil
}

func (r *ManagerRepositoryPostgres) UpdateManager(id int, wantManager internal_types.Manager) error {
	var err error
	updatePersonQuery := `
	UPDATE persons
	SET email=$1, phone=$2, address=$3
	WHERE id IN (SELECT person_id FROM managers WHERE id=$4)`

	updateWorkInfoQuery := `
	UPDATE manager_work_info
	SET bank_account=$1, capital_managment=$2, profit_percent_day=$3
	WHERE id IN (SELECT work_info_id FROM managers WHERE managers.id=$4)
`

	_, err = r.db.Exec(updatePersonQuery, wantManager.Email, wantManager.Phone, wantManager.Address, id)
	if err != nil {
		return err
	}
	_, err = r.db.Exec(updateWorkInfoQuery, wantManager.BankAccount, wantManager.CapitalManagment, wantManager.ProfitPercentDay, id)
	if err != nil {
		return err
	}

	return err
}
